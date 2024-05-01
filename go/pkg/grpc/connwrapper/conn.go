package connwrapper

import (
	"context"
	go_errors "errors"
	"io"
	"log"
	"net"
	"prodvana/errlog"
	"prodvana/perrors"
	"strings"

	"github.com/pkg/errors"
	grpc_codes "google.golang.org/grpc/codes"
)

type wrapperOptions struct {
	logger *log.Logger
}

type WrapperOptions func(wrapperOptions) wrapperOptions

func WithLogger(logger *log.Logger) WrapperOptions {
	return func(o wrapperOptions) wrapperOptions {
		o.logger = logger
		return o
	}
}

type BlobRequest interface {
	GetBlob() []byte
}

type StreamServer[RecvT BlobRequest, SendT any] interface {
	Recv() (RecvT, error)
	Send(SendT) error
}

func shouldReport(err error) bool {
	if strings.Contains(err.Error(), "timed out getting a new connection to Kube API") {
		// avoid double logging as this is already logged on agent interaction
		return false
	}
	if go_errors.Is(err, io.EOF) {
		return false
	}
	if go_errors.Is(err, io.ErrClosedPipe) {
		return false
	}
	if go_errors.Is(err, context.Canceled) {
		return false
	}
	if go_errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	switch perrors.GrpcErrCode(err) {
	case grpc_codes.Canceled, grpc_codes.DeadlineExceeded, grpc_codes.Internal:
		return false
	}
	return true
}

// Connect a golang net.Conn interface with a grpc bi-directional streaming API.
// This function blocks until the connection is terminated or errored out.
// The caller is responsible for making sure that after this function call,
// with the exception of startWrite, the bi-directional stream only sends and
// receives bytes, not any other data types.
// Caller is responsible for closing both conn and strm.
func ConnectConnAndStreamingServer[RecvT BlobRequest, SendT any](conn net.Conn, strm StreamServer[RecvT, SendT], sendFactory func([]byte) SendT, startWrite <-chan struct{}, options ...WrapperOptions) {
	compiledOptions := wrapperOptions{}
	for _, o := range options {
		compiledOptions = o(compiledOptions)
	}
	readClosed := make(chan struct{})
	writeClosed := make(chan struct{})

	go func() {
		defer func() {
			close(readClosed)
		}()

		select {
		case <-startWrite:
		case <-writeClosed:
			return
		}

		buf := make([]byte, 4096)
		for {
			var n int
			var err error
			readDone := make(chan struct{})
			go func() {
				defer close(readDone)
				n, err = conn.Read(buf)
			}()
			select {
			case <-writeClosed:
				return
			case <-readDone:
			}

			if err != nil {
				if shouldReport(err) {
					errlog.ReportError(context.TODO(), errors.Wrap(err, "failed to read from conn"), errlog.LevelWarning)
				}
				return
			}

			sendReq := sendFactory(buf[:n])
			if err := strm.Send(sendReq); err != nil {
				if shouldReport(err) {
					errlog.ReportError(context.TODO(), errors.Wrap(err, "failed to send to stream"), errlog.LevelWarning)
				}
				return
			}
		}
	}()

	go func() {
		defer func() {
			close(writeClosed)
		}()

		for {
			var req RecvT
			var err error
			readDone := make(chan struct{})

			go func() {
				defer close(readDone)
				req, err = strm.Recv()
			}()

			select {
			case <-readClosed:
				return
			case <-readDone:
			}

			if err != nil {
				if shouldReport(err) {
					errlog.ReportError(context.TODO(), errors.Wrap(err, "failed to read from stream"), errlog.LevelWarning)
				}
				return
			}
			_, err = conn.Write(req.GetBlob())
			if err != nil {
				if shouldReport(err) {
					errlog.ReportError(context.TODO(), errors.Wrap(err, "failed to send to conn"), errlog.LevelWarning)
				}
				return
			}
		}
	}()

	<-readClosed
	<-writeClosed
}
