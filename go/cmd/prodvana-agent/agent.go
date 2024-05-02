package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/prodvana/prodvana-public/go/pkg/client"
	"github.com/prodvana/prodvana-public/go/pkg/counter"
	"github.com/prodvana/prodvana-public/go/pkg/grpc/connwrapper"
	"github.com/prodvana/prodvana-public/go/pkg/timeutils"
	"github.com/prodvana/prodvana-public/go/pkg/types"

	gcp_metadata "cloud.google.com/go/compute/metadata"
	client_auth "github.com/prodvana/prodvana-public/go/pkg/auth/client"
	agent_pb "github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/agent"
	"github.com/prodvana/prodvana-public/go/prodvana-sdk/proto/prodvana/environment"

	"github.com/pkg/errors"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/proxy"
)

const (
	heartbeatInterval        = time.Minute
	heartbeatFailureInterval = 10 * time.Second
	debugInterval            = 5 * time.Minute
	startInformerTimeout     = 10 * time.Second
	reportMetadataInterval   = 5 * time.Minute
	jitterPercent            = 10
	failureCounterWindow     = 2 * time.Minute

	poolsCount                           = 20
	poolSize                             = 10
	proxyConnectionInitializationTimeout = 30 * time.Second
)

type lastSuccessTime struct {
	mu    sync.Mutex
	value time.Time
}

func (l *lastSuccessTime) Now() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.value = time.Now()
}

func (l *lastSuccessTime) Get() time.Time {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.value
}

func heartbeatForever(ctx context.Context, client agent_pb.AgentInteractionClient, clusterId types.RuntimeId, lastSuccessTime *lastSuccessTime) {
	for {
		_, err := client.Heartbeat(ctx, &agent_pb.HeartbeatReq{
			ClusterId: string(clusterId),
		})
		if err != nil {
			log.Printf("Heartbeat failed: %+v", err)
			select {
			case <-time.After(timeutils.JitterPercentPositive(heartbeatFailureInterval, jitterPercent)):
			case <-ctx.Done():
				return
			}
		} else {
			lastSuccessTime.Now()
			log.Printf("Heartbeat succeeded")
			select {
			case <-time.After(timeutils.JitterPercentPositive(heartbeatInterval, jitterPercent)):
			case <-ctx.Done():
				return
			}
		}
	}
}

func gatherClusterMetadata(ctx context.Context) (*environment.ClusterMetadata, error) {
	if gcp_metadata.OnGCE() {
		gcpProjectId, err := gcp_metadata.ProjectID()
		if err != nil {
			return nil, err
		}

		// Magic strings from https://cloud.google.com/kubernetes-engine/docs/concepts/workload-identity#instance_attributes
		gcpLocation, err := gcp_metadata.InstanceAttributeValue("cluster-location")
		if err != nil {
			return nil, err
		}
		gcpClusterName, err := gcp_metadata.InstanceAttributeValue("cluster-name")
		if err != nil {
			return nil, err
		}

		return &environment.ClusterMetadata{
			ClusterMetadata: &environment.ClusterMetadata_Gke{
				Gke: &environment.GKEClusterMetadata{
					ProjectName:     gcpProjectId,
					ClusterLocation: gcpLocation,
					ClusterName:     gcpClusterName,
				},
			},
		}, nil
	}

	return nil, nil
}

func reportClusterMetadata(ctx context.Context, client agent_pb.AgentInteractionClient, clusterId types.RuntimeId) {
	for {
		clusterMetadata, err := gatherClusterMetadata(ctx)
		log.Printf("Reporting cluster metadata %+v with error %+v", clusterMetadata, err)

		var errStr string
		if err != nil {
			errStr = err.Error()
		}

		if _, err := client.ReportClusterMetadata(ctx, &agent_pb.ReportClusterMetadataReq{
			RuntimeId:       string(clusterId),
			ClusterMetadata: clusterMetadata,
			Error:           errStr,
		}); err != nil {
			log.Printf("Failed to report cluster metadata: %+v", err)
		}

		select {
		case <-time.After(timeutils.JitterPercent(reportMetadataInterval, jitterPercent)):
		case <-ctx.Done():
			return
		}
	}
}

type conns struct {
	nextId      int
	outstanding int
	inUse       int
	lock        sync.Mutex
	cond        *sync.Cond
}

func (c *conns) newInuse() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.inUse++
	c.cond.Signal()
}

func (c *conns) closeInuse() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.inUse--
	c.cond.Signal()
}

func (c *conns) closeOutstanding() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.outstanding--
	c.cond.Signal()
}

func (c *conns) id() int {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.nextId++
	c.cond.Signal()

	return c.nextId
}

func proxyAPIServer(ctx context.Context, logger *log.Logger, c *conns, client agent_pb.AgentInteractionClient, cfg *rest.Config, runtimeId types.RuntimeId, connId string, connInitFailureCounter *counter.RollingWindowCounter) error {
	if ctx.Err() != nil {
		return errors.Wrap(ctx.Err(), "ctx closed: proxyAPIServer shutting down")
	}

	logger.Println("Issuing ProxyAPIServer RPC")
	start := time.Now()
	successChan := make(chan struct{})
	wrappedCtx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-time.After(proxyConnectionInitializationTimeout):
			logger.Printf("WARNING: Server took a long time to return ProxyAPIServer stream: %s\n", time.Since(start))
			cancel()
		case <-successChan:
			return
		}
	}()
	strm, err := client.ProxyAPIServer(wrappedCtx)
	logger.Println("ProxyAPIServer RPC returned")
	if err != nil {
		connInitFailureCounter.Add()
		return errors.Wrap(err, "initializing stream to pvn apiserver failed")
	}
	successChan <- struct{}{}

	if err := strm.Send(&agent_pb.ProxyAPIServerReq{
		Msg: &agent_pb.ProxyAPIServerReq_RuntimeId{
			RuntimeId: string(runtimeId),
		},
		ConnId: connId,
	}); err != nil {
		connInitFailureCounter.Add()
		return errors.Wrap(err, "error sending initial header")
	}

	// We have a potentially live connection - make sure it is closed before returning
	defer func() {
		_ = strm.CloseSend()
	}()

	logger.Println("Waiting for header from pvn apiserver")
	recv, err := strm.Recv()
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return errors.Wrap(err, "error reading initial header")
	}
	logger.Println("Processing header from pvn apiserver")

	c.newInuse()
	defer c.closeInuse()

	switch recv.Msg.(type) {
	case *agent_pb.ProxyAPIServerResp_Type_:
		typ := recv.Msg.(*agent_pb.ProxyAPIServerResp_Type_).Type
		if typ != agent_pb.ProxyAPIServerResp_PROXY_APISERVER {
			return fmt.Errorf("agent only supports proxying apiserver - unknown connection type %s", typ)
		}
	default:
		return fmt.Errorf("Expected header message to contain request type, but it contained %T instead", recv.Msg)
	}

	// TODO: Use unix socket instead?
	logger.Println("Launching APIserver proxy")
	server, err := proxy.NewServer("", "/", "", nil /* TODO: Disabling filters is considered dangerous */, cfg, time.Minute, false)
	if err != nil {
		return errors.Wrap(err, "error launching apiserver proxy")
	}

	l, err := server.Listen("localhost", 0)
	if err != nil {
		return err
	}

	// Listener is running. Shut it down before returning
	defer l.Close()

	logger.Println("Listening on", l.Addr())

	go func() {
		logger.Println("Serving on listener")
		if err := server.ServeOnListener(l); err != nil {
			logger.Println("Exiting server", err)
		}
	}()

	logger.Println("Connecting to proxy")
	conn, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		return errors.Wrap(err, "error connecting to proxied apiserver")
	}
	defer conn.Close()

	logger.Println("Starting proxy loop, waiting for interaction to finish")
	startWrite := make(chan struct{})
	close(startWrite) // we don't have any dependencies before writes can start
	connwrapper.ConnectConnAndStreamingServer[*agent_pb.ProxyAPIServerResp, *agent_pb.ProxyAPIServerReq](conn, strm, func(blob []byte) *agent_pb.ProxyAPIServerReq {
		return &agent_pb.ProxyAPIServerReq{
			Msg: &agent_pb.ProxyAPIServerReq_Blob{
				Blob: blob,
			},
		}
	}, startWrite, connwrapper.WithLogger(logger))

	logger.Println("Closing stream")

	return strm.CloseSend()
}

func proxyAPIServerForever(ctx context.Context, poolId int, count int, client agent_pb.AgentInteractionClient, cfg *rest.Config, runtimeId types.RuntimeId, connInitFailureCounter *counter.RollingWindowCounter) {
	c := &conns{}
	c.cond = sync.NewCond(&c.lock)

	newConn := func() {
		defer c.closeOutstanding()
		idx := c.id()
		connId := fmt.Sprintf("pool %d, conn %d, rtm %s", poolId, idx, runtimeId)
		// logger := log.New(os.Stderr, fmt.Sprintf("[%s] ", connId), log.Lshortfile|log.Lmicroseconds)
		// do not log by default, it's very spammy
		logger := log.New(io.Discard, "", 0)
		if err := proxyAPIServer(ctx, logger, c, client, cfg, runtimeId, connId, connInitFailureCounter); err != nil {
			logger.Printf("Error proxying %s", err)
		}
	}

	// logger := log.New(os.Stderr, fmt.Sprintf("[pool %2d] ", poolId), log.Lshortfile|log.Lmicroseconds)
	logger := log.New(io.Discard, fmt.Sprintf("[pool %2d] ", poolId), log.Lshortfile|log.Lmicroseconds)

	c.outstanding++
	go newConn()

	for {
		c.lock.Lock()

		for c.outstanding-c.inUse < count {
			logger.Printf("Outstanding: %d, in use: %d, target: %d - creating new connection", c.outstanding, c.inUse, count)
			c.outstanding++
			go newConn()
		}

		c.cond.Wait()
		c.lock.Unlock()
		time.Sleep(time.Millisecond)
	}
}

func main() {
	port := flag.Int("port", 5100, "Port to run http server on")
	apiserverAddr := flag.String("server-addr", "apiserver:5000", "Address for Prodvana upstream server, to run agent interaction commands.")
	skipTls := flag.Bool("skip-tls", false, "Skip TLS verification for the connection to the agent interaction.")
	token := flag.String("auth", "", "API Token for authenticating with Prodvana.")
	clusterIdStr := flag.String("clusterid", "", "Identifier for the cluster the agent is running in.")
	// TODO(mike): remove once apiserver is rolled out and agents updated to stop passing this flag
	flag.Bool("disable-flagger", true, "Deprecated, no longer used.")
	flag.Bool("enable-rollouts", false, "Deprecated, no longer used.")
	flag.Parse()
	if err := client.SetDefaultApiServerAddr(*apiserverAddr, *skipTls); err != nil {
		log.Fatal(err)
	}

	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	var clusterId types.RuntimeId
	if *clusterIdStr == "" {
		log.Fatal("-clusterid not passed")
	} else {
		clusterId = types.RuntimeId(*clusterIdStr)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	authToken := client_auth.AuthToken{Token: *token}
	log.Printf("Running agent with token: %s\n", *token)
	conn := client.GetApiserverConnWithAuth(&authToken)
	defer conn.Close()
	primaryClient := agent_pb.NewAgentInteractionClient(conn)

	k8sConfig, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	stopCh := make(chan struct{})
	defer func() {
		close(stopCh)
	}()

	primaryClientLastSuccessTime := &lastSuccessTime{}

	go heartbeatForever(ctx, primaryClient, clusterId, primaryClientLastSuccessTime)
	go reportClusterMetadata(ctx, primaryClient, clusterId)

	// proxy apiserver via creating dedicated clients, required to get around GCP load balancer's connection concurrency limit of 100.
	// (this is a poor man's connection pool)
	proxyFailuresByConn := make([]*counter.RollingWindowCounter, poolsCount)
	for poolId := 0; poolId < poolsCount; poolId++ {
		conn, err := client.MakeApiServerConnWithAuth(client.ApiServerAddr, &authToken, *skipTls)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		client := agent_pb.NewAgentInteractionClient(conn)
		failureCounter := counter.NewRollingWindowCounter(failureCounterWindow, 1*time.Second)
		proxyFailuresByConn[poolId] = failureCounter
		go proxyAPIServerForever(ctx, poolId, poolSize, client, k8sConfig, clusterId, failureCounter)
	}

	log.Printf("Agent running on :%d", *port)

	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		failure := false
		messages := ""
		heartbeatSuccessTime := primaryClientLastSuccessTime.Get()
		if time.Since(heartbeatSuccessTime) > heartbeatInterval*2 {
			failure = true
			messages += fmt.Sprintf("heartbeat has not succeeded since %s\n", heartbeatSuccessTime)
		} else {
			messages += fmt.Sprintf("heartbeat last succeeded at %s\n", heartbeatSuccessTime)
		}
		for poolId, failureCounter := range proxyFailuresByConn {
			count := failureCounter.Count()
			if count > (poolSize * 2) {
				failure = true
			}
			messages += fmt.Sprintf("proxy pool %d has %d connection initialization failures in the last %s\n", poolId, count, failureCounterWindow)
		}

		if failure {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "failures:\n%s\n", messages)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "ok:\n%s\n", messages)
		}
	})
	_ = http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
