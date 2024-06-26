// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: prodvana/billing/manager.proto

/*
Package billing is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package billing

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_BillingManager_GetUsageMetrics_0(ctx context.Context, marshaler runtime.Marshaler, client BillingManagerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetUsageMetricsReq
	var metadata runtime.ServerMetadata

	msg, err := client.GetUsageMetrics(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BillingManager_GetUsageMetrics_0(ctx context.Context, marshaler runtime.Marshaler, server BillingManagerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetUsageMetricsReq
	var metadata runtime.ServerMetadata

	msg, err := server.GetUsageMetrics(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBillingManagerHandlerServer registers the http handlers for service BillingManager to "mux".
// UnaryRPC     :call BillingManagerServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBillingManagerHandlerFromEndpoint instead.
func RegisterBillingManagerHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BillingManagerServer) error {

	mux.Handle("GET", pattern_BillingManager_GetUsageMetrics_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/prodvana.billing.BillingManager/GetUsageMetrics", runtime.WithHTTPPathPattern("/v1/billing/usage"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BillingManager_GetUsageMetrics_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BillingManager_GetUsageMetrics_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBillingManagerHandlerFromEndpoint is same as RegisterBillingManagerHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBillingManagerHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBillingManagerHandler(ctx, mux, conn)
}

// RegisterBillingManagerHandler registers the http handlers for service BillingManager to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBillingManagerHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBillingManagerHandlerClient(ctx, mux, NewBillingManagerClient(conn))
}

// RegisterBillingManagerHandlerClient registers the http handlers for service BillingManager
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BillingManagerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BillingManagerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BillingManagerClient" to call the correct interceptors.
func RegisterBillingManagerHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BillingManagerClient) error {

	mux.Handle("GET", pattern_BillingManager_GetUsageMetrics_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/prodvana.billing.BillingManager/GetUsageMetrics", runtime.WithHTTPPathPattern("/v1/billing/usage"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BillingManager_GetUsageMetrics_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BillingManager_GetUsageMetrics_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BillingManager_GetUsageMetrics_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "billing", "usage"}, ""))
)

var (
	forward_BillingManager_GetUsageMetrics_0 = runtime.ForwardResponseMessage
)
