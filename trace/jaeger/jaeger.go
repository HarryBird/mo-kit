package jaeger

import (
	"errors"

	stdjaeger "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var (
	ErrInvalidCollectorEndpoint  = errors.New("mo-kit -> jaeger: invalid collector endpoint")
	ErrInvalidProviderServiceKey = errors.New("mo-kit -> jaeger: invalid provider service key")
)

func DefaultExporter(col ICollector) (*stdjaeger.Exporter, error) {
	if col.GetEndpoint() == "" {
		return nil, ErrInvalidCollectorEndpoint
	}

	opts := []stdjaeger.CollectorEndpointOption{
		stdjaeger.WithEndpoint(col.GetEndpoint()),
	}

	if col.GetUsername() != "" {
		opts = append(opts, stdjaeger.WithUsername(col.GetUsername()))
	}

	if col.GetPassword() != "" {
		opts = append(opts, stdjaeger.WithPassword(col.GetPassword()))
	}

	return stdjaeger.New(stdjaeger.WithCollectorEndpoint(opts...))
}

func DefaultProvider(key string, col ICollector) (*tracesdk.TracerProvider, error) {
	if key == "" {
		return nil, ErrInvalidProviderServiceKey
	}

	exp, err := DefaultExporter(col)
	if err != nil {
		return nil, err
	}

	return tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(key),
		)),
	), nil
}
