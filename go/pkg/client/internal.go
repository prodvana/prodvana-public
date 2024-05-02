package client

import (
	dd_internal_pb "prodvana/proto/prodvana_internal/dynamic_delivery"
	"sync"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	dynamicDeliveryAddr = "dynamic-delivery:5050"
)

var (
	dynamicDeliveryManagerClient     dd_internal_pb.DynamicDeliveryManagerClient
	dynamicDeliveryManagerClientOnce sync.Once
)

func GetDefaultDynamicDeliveryManagerClient() dd_internal_pb.DynamicDeliveryManagerClient {
	dynamicDeliveryManagerClientOnce.Do(func() {
		conn, err := grpc.NewClient(dynamicDeliveryAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
			grpc.WithDefaultCallOptions(
				grpc.WaitForReady(true),
				grpc.MaxCallRecvMsgSize(200*1024*1024), // 200MB. Allow APIserver to read larger responses from DD
			),
		)
		if err != nil {
			panic(err)
		}
		dynamicDeliveryManagerClient = dd_internal_pb.NewDynamicDeliveryManagerClient(conn)
	})
	return dynamicDeliveryManagerClient

}
