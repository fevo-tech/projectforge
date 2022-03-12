// Content managed by Project Forge, see [projectforge.md] for details.
package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"projectforge.dev/projectforge/app/util"
)

var (
	initialized    = false
	tracerProvider *sdktrace.TracerProvider
)

func InitializeIfNeeded(enabled bool, logger *zap.SugaredLogger) bool {
	if initialized {
		return false
	}
	if enabled {
		Initialize(logger)
	}
	return true
}

func Initialize(logger *zap.SugaredLogger) {
	if initialized {
		logger.Warn("double telemetry initialization")
		return
	}
	otel.SetErrorHandler(&ErrHandler{logger: logger})
	initialized = true

	endpoint := "localhost:55681"
	if env := util.GetEnv("telemetry_endpoint"); env != "" {
		endpoint = env
	}
	logger.Debugf("initializing OpenTelemetry tracing using endpoint [%s]", endpoint)
	tp, err := buildTP(endpoint)
	if err != nil {
		logger.Errorf("unable to create tracing provider: %+v", err)
		return
	}
	tracerProvider = tp
	p := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(p)
}

func buildTP(endpoint string) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint(endpoint), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(util.AppKey))),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func Close() error {
	return tracerProvider.Shutdown(context.Background())
}

func StartSpan(ctx context.Context, spanName string, logger *zap.SugaredLogger, opts ...interface{}) (context.Context, *Span, *zap.SugaredLogger) {
	return spanCreate(ctx, spanName, logger, opts...)
}

func StartAsyncSpan(ctx context.Context, spanName string, logger *zap.SugaredLogger, opts ...interface{}) (context.Context, *Span, *zap.SugaredLogger) {
	parentSpan := trace.SpanFromContext(ctx)
	asyncChildCtx := trace.ContextWithSpan(context.Background(), parentSpan)
	return spanCreate(asyncChildCtx, spanName, logger, opts...)
}

func spanCreate(ctx context.Context, spanName string, logger *zap.SugaredLogger, opts ...interface{}) (context.Context, *Span, *zap.SugaredLogger) {
	tr := otel.GetTracerProvider().Tracer(util.AppKey)
	ssos := []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindServer)}
	for _, opt := range opts {
		o, ok := opt.(trace.SpanStartOption)
		if ok {
			ssos = append(ssos, o)
		}
	}
	ctx, ot := tr.Start(ctx, spanName, ssos...)
	sp := &Span{OT: ot}
	if logger != nil {
		logger = LoggerFor(logger, sp)
	}
	return ctx, sp, logger
}
