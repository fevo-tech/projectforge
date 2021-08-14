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

	"{{{ .Package }}}/app/util"
)

type Service struct {
	enabled bool
	tp      *sdktrace.TracerProvider
	logger  *zap.SugaredLogger
}

func NewService(logger *zap.SugaredLogger) *Service {
	tp, err := buildTP("localhost:55681")
	if err != nil {
		return &Service{enabled: false, logger: logger}
	}

	p := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(p)

	return &Service{enabled: true, tp: tp, logger: logger}
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

func (s *Service) Close() error {
	return s.tp.Shutdown(context.Background())
}

func (s *Service) StartSpan(ctx context.Context, tracerKey string, spanName string) (context.Context, trace.Span) {
	tr := otel.GetTracerProvider().Tracer(tracerKey)
	return tr.Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindServer))
}
