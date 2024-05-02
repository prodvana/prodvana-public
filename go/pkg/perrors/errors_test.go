package perrors

import (
	"context"
	"testing"

	"github.com/prodvana/prodvana-public/go/pkg/errlog"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
)

func TestGrpcForwarding(t *testing.T) {
	require.NotEqual(t, grpc_codes.NotFound, grpc_status.Code(ConvertGrpcErrorInServerInterceptor(context.Background(), grpc_status.Error(grpc_codes.NotFound, "test"))))
	require.NotEqual(t, grpc_codes.NotFound, grpc_status.Code(ConvertGrpcErrorInServerInterceptor(context.Background(), errors.Wrap(grpc_status.Error(grpc_codes.NotFound, "test"), "wrapped"))))
	require.Equal(t, grpc_codes.NotFound, grpc_status.Code(ConvertGrpcErrorInServerInterceptor(context.Background(), GrpcError(grpc_codes.NotFound, "test"))))
	require.Equal(t, grpc_codes.NotFound, grpc_status.Code(ConvertGrpcErrorInServerInterceptor(context.Background(), errors.Wrap(GrpcError(grpc_codes.NotFound, "test"), "wrapped"))))
	require.Equal(t, grpc_codes.OK, grpc_status.Code(ConvertGrpcErrorInServerInterceptor(context.Background(), nil)))
}

func TestGrpcForwardingRespectsSkip(t *testing.T) {
	skippedUnforwarded := errlog.Skip(grpc_status.Error(grpc_codes.Internal, "skipped unforwarded"))
	require.True(t, errlog.IsSkipped(skippedUnforwarded))

	skippedForwarded := errlog.Skip(GrpcError(grpc_codes.Internal, "skipped forwarded"))
	require.True(t, errlog.IsSkipped(ConvertGrpcErrorInServerInterceptor(context.Background(), skippedForwarded)))
}
