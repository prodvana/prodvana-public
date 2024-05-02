package perrors

import (
	"context"
	go_errors "errors"
	"fmt"
	"syscall"

	"github.com/prodvana/prodvana-public/go/pkg/errlog"

	"cloud.google.com/go/spanner"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
)

func UserFacingInternalGrpcErrorf(eventID *sentry.EventID, format string, a ...interface{}) error {
	base := fmt.Sprintf(format, a...)
	var eventStr string
	if eventID != nil {
		eventStr = fmt.Sprintf("\n(Event ID: %s)", string(*eventID))
	}
	return GrpcErrorf(grpc_codes.Internal, "%s\nPlease try again. If this error persists, please contact support@prodvana.io.%s", base, eventStr)
}

// similar to spanner.ErrCode but handles wrapped errors
func GrpcErrCode(err error) grpc_codes.Code {
	var grpcErr interface {
		error
		GRPCStatus() *grpc_status.Status
	}
	if go_errors.As(err, &grpcErr) {
		return spanner.ErrCode(grpcErr)
	}
	var innerErr *grpcError
	if go_errors.As(err, &innerErr) {
		return spanner.ErrCode(innerErr.InnerGrpcError())
	}
	return spanner.ErrCode(err)
}

func AwsErrCode(err error) string {
	var awsErr awserr.Error
	if go_errors.As(err, &awsErr) {
		return awsErr.Code()
	}
	return ""
}

// detect transient errors, e.g. https://stackoverflow.com/questions/45793299/how-to-check-if-error-is-tls-handshake-timeout-in-go
func IsTransientNetworkError(err error) bool {
	if go_errors.Is(err, syscall.ECONNRESET) || go_errors.Is(err, syscall.EPIPE) {
		return true
	}
	type timeoutErr interface {
		Timeout() bool
	}
	var tErr timeoutErr
	if go_errors.As(err, &tErr) && tErr.Timeout() {
		return true
	}
	type temporaryErr interface {
		Temporary() bool
	}
	var teErr temporaryErr
	if go_errors.As(err, &teErr) && teErr.Temporary() {
		return true
	}
	return false
}

func IsDeadlineExceeded(err error) bool {
	if go_errors.Is(err, context.DeadlineExceeded) {
		return true
	}
	return GrpcErrCode(err) == grpc_codes.DeadlineExceeded
	// var innerErr *grpcError
	// if go_errors.As(err, &innerErr) {

	// 	if nativeGrpcErr, ok := innerErr.InnerGrpcError().(interface {
	// 		error
	// 		GRPCStatus() *grpc_status.Status
	// 	}); ok {
	// 		return nativeGrpcErr.GRPCStatus().Code() == grpc_codes.DeadlineExceeded
	// 	}
	// }
	// return false
}

// wrapper for grpc errors to avoid accidentally forwarding grpc errors from spanner and other backends to the user.
type grpcError struct {
	grpcErr error
}

// Error implements error
func (e *grpcError) Error() string {
	return e.grpcErr.Error()
}

func GrpcErrorf(c grpc_codes.Code, format string, a ...interface{}) error {
	return &grpcError{
		grpcErr: grpc_status.Errorf(c, format, a...),
	}
}

func GrpcError(c grpc_codes.Code, msg string) error {
	return &grpcError{
		grpcErr: grpc_status.Error(c, msg),
	}
}

func (e *grpcError) InnerGrpcError() error {
	return e.grpcErr
}

func (e *grpcError) Unwrap() error {
	return e.grpcErr
}

// Takes an error and return back an appropriate error that should be passed
// to grpc server handler. The goal here is to only pass grpc status set by
// perrors.grpcError and not native grpc status package, to avoid accidentally
// a grpc error from a dependency to client.
func ConvertGrpcErrorInServerInterceptor(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if ctx.Err() != nil {
		// when context has an error, the underlying error is not trustworthy
		return ctx.Err()
	}
	var resultErr error
	var innerErr *grpcError
	if go_errors.As(err, &innerErr) {
		resultErr = innerErr.InnerGrpcError()
	} else if _, ok := grpc_status.FromError(err); ok {
		// wrap to avoid forwarding. GRPC Server does not unwrap on its own.
		resultErr = errors.Errorf("unforwarded grpc error: %+v", err)
	} else {
		return err
	}
	// check if this error should not be reported to the errlog
	if errlog.IsSkipped(err) {
		return errlog.Skip(resultErr)
	}
	return resultErr
}
