// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: prodvana/blobs/blobs_manager.proto

/*
Package blobs is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package blobs

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

func request_BlobsManager_GetCasBlob_0(ctx context.Context, marshaler runtime.Marshaler, client BlobsManagerClient, req *http.Request, pathParams map[string]string) (BlobsManager_GetCasBlobClient, runtime.ServerMetadata, error) {
	var protoReq GetCasBlobReq
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	stream, err := client.GetCasBlob(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterBlobsManagerHandlerServer registers the http handlers for service BlobsManager to "mux".
// UnaryRPC     :call BlobsManagerServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBlobsManagerHandlerFromEndpoint instead.
func RegisterBlobsManagerHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BlobsManagerServer) error {

	mux.Handle("GET", pattern_BlobsManager_GetCasBlob_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterBlobsManagerHandlerFromEndpoint is same as RegisterBlobsManagerHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBlobsManagerHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterBlobsManagerHandler(ctx, mux, conn)
}

// RegisterBlobsManagerHandler registers the http handlers for service BlobsManager to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBlobsManagerHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBlobsManagerHandlerClient(ctx, mux, NewBlobsManagerClient(conn))
}

// RegisterBlobsManagerHandlerClient registers the http handlers for service BlobsManager
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BlobsManagerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BlobsManagerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BlobsManagerClient" to call the correct interceptors.
func RegisterBlobsManagerHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BlobsManagerClient) error {

	mux.Handle("GET", pattern_BlobsManager_GetCasBlob_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/prodvana.blobs.BlobsManager/GetCasBlob", runtime.WithHTTPPathPattern("/v1/blobs/cas/{id=*}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BlobsManager_GetCasBlob_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BlobsManager_GetCasBlob_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BlobsManager_GetCasBlob_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "blobs", "cas", "id"}, ""))
)

var (
	forward_BlobsManager_GetCasBlob_0 = runtime.ForwardResponseStream
)