package errlog

import (
	"context"
	go_errors "errors"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"
)

type Level sentry.Level

var (
	LevelDebug   = Level(sentry.LevelDebug)
	LevelInfo    = Level(sentry.LevelInfo)
	LevelWarning = Level(sentry.LevelWarning)
	LevelError   = Level(sentry.LevelError)
	LevelFatal   = Level(sentry.LevelFatal)
)

type skipError struct {
	error
}

// return an error that will be skipped by errlog reporting
func Skip(err error) error {
	return &skipError{err}
}

func (e *skipError) Unwrap() error {
	return e.error
}

func IsSkipped(err error) bool {
	if err == nil {
		return false
	}
	var serr *skipError
	return go_errors.As(err, &serr)
}

func Init() error {
	sentryDsn := os.Getenv("SENTRY_DSN")
	if sentryDsn == "" {
		log.Printf("WARNING: Skipping sentry initialization, SENTRY_DSN not set.")
		return nil
	}
	gitHash := os.Getenv("GIT_HASH")
	if gitHash == "" {
		log.Printf("WARNING: GIT_HASH not set, sentry release will not be set")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		Release:          gitHash,
		AttachStacktrace: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to init sentry")
	}
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		environment := os.Getenv("PVN_ENVIRONMENT")
		if environment == "" {
			environment = "unknown"
		}
		scope.SetTag("environment", environment)
	})
	return nil
}

func Flush() {
	flushed := sentry.Flush(5 * time.Second)
	if !flushed {
		log.Printf("WARNING: Failed to flush sentry")
	}
}

func currentHub(ctx context.Context) *sentry.Hub {
	hub := sentry.GetHubFromContext(ctx)
	if hub == nil {
		hub = sentry.CurrentHub()
	}
	return hub
}

// See: https://docs.honeycomb.io/working-with-your-data/direct-trace-links/
func honeycombTraceURL(ctx context.Context) (bool, string) {
	spanContext := trace.SpanContextFromContext(ctx)
	if !spanContext.IsValid() {
		return false, ""
	}
	environment := os.Getenv("PVN_ENVIRONMENT")
	switch environment {
	case "prod", "production":
		environment = "production"
	case "staging":
		environment = "staging"
	default:
		return false, ""
	}
	dataset := os.Getenv("PVN_SERVICE")
	if dataset == "" {
		return false, ""
	}
	traceID := spanContext.TraceID()
	// TODO(mike): I couldn't figure out how to extract the start ts from the span,
	// so using now - 30m to increase our chances of the trace falling in the timeframe
	start := time.Now().Add(-30 * time.Minute)
	end := start.Add(2 * time.Hour)

	return true, fmt.Sprintf(
		// 1 == Environment
		// 2 == Dataset
		// 3 == TraceID
		// 4 == Start Timestamp
		// 5 == End Timestamp
		"https://ui.honeycomb.io/prodvana/environments/%[1]s/datasets/%[2]s/trace?trace_id=%[3]s&trace_start_ts=%[4]d&trace_end_ts=%[5]d",
		environment,      // 1
		dataset,          // 2
		traceID.String(), // 3
		start.Unix(),     // 4
		end.Unix(),       // 5
	)
}

func annotateTrace(ctx context.Context, eventID *sentry.EventID) {
	if eventID == nil {
		return
	}
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attribute.Key("sentry_event_id").String(string(*eventID)))
}

func ReportError(ctx context.Context, err error, level Level) *sentry.EventID {
	if IsSkipped(err) {
		return nil
	}
	log.Printf("%+v", err)
	hub := currentHub(ctx)
	var eventID *sentry.EventID
	hub.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.Level(level))
		if ok, honeycombURL := honeycombTraceURL(ctx); ok {
			scope.SetTag("honeycomb_trace_url", honeycombURL)

		}
		eventID = hub.CaptureException(err)
		annotateTrace(ctx, eventID)
	})
	return eventID
}

func ReportErrorWithRateLimit(ctx context.Context, err error, level Level, rateLimiter *rate.Limiter) *sentry.EventID {
	if rateLimiter != nil && rateLimiter.Allow() {
		return ReportError(ctx, err, level)
	}
	return nil
}

func ReportErrorAndCrash(ctx context.Context, err error) {
	hub := currentHub(ctx)
	hub.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.Level(LevelFatal))
		hub.CaptureException(err)
	})
	Flush()
	log.Fatalf("%+v", err)
}

func FatalErr(err error) {
	ReportErrorAndCrash(context.Background(), err)
}

func FatalErrf(err error, msg string, args ...interface{}) {
	ReportErrorAndCrash(context.Background(), errors.Wrapf(err, msg, args...))
}

func Fatal(msg interface{}) {
	ReportErrorAndCrash(context.Background(), errors.Errorf("%s", msg))
}

func Fatalf(msg string, args ...interface{}) {
	ReportErrorAndCrash(context.Background(), errors.Errorf(msg, args...))
}

func WrapCtx(ctx context.Context) context.Context {
	hub := currentHub(ctx).Clone()
	ctx = sentry.SetHubOnContext(ctx, hub)
	return ctx
}

func SetTag(ctx context.Context, key, value string) {
	hub := currentHub(ctx)
	hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag(key, value)
	})
}

func SetContext(ctx context.Context, sectionTitle string, values map[string]interface{}) {
	hub := currentHub(ctx)
	hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetContext(sectionTitle, values)
	})
}

func RecoverAndCrash() {
	err := recover()
	if err != nil {
		sentry.CurrentHub().Recover(err)
		Flush()
		log.Fatalf("%+v\n%s", err, debug.Stack())
	}
}
