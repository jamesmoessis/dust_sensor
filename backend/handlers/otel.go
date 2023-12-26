package handlers

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"

	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// auth is configured by env var:
// OTEL_EXPORTER_OTLP_METRICS_HEADERS: 'Lightstep-Access-Token: ***'
const lightStepEndpoint = "ingest.lightstep.com"
const lightStepPath = "/metrics/otlp/v0.9"

func NewResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes("",
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}

func NewExporter(ctx context.Context) (metric.Exporter, error) {
	return otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(lightStepEndpoint),
		otlpmetrichttp.WithURLPath(lightStepPath),
		otlpmetrichttp.WithInsecure(),
	)
}

type Recorder struct {
	Exporter metric.Exporter
	Resource resource.Resource
}

// todo send the values together
func (r *Recorder) RecordIntMetric(name string, val int64) error {
	return r.Exporter.Export(context.Background(), &metricdata.ResourceMetrics{
		Resource: &r.Resource,
		ScopeMetrics: []metricdata.ScopeMetrics{
			{
				Scope: instrumentation.Scope{
					Name:    "dust_sensor_api",
					Version: "v1",
				},
				Metrics: []metricdata.Metrics{
					{
						Name: name,
						Data: metricdata.Gauge[int64]{
							DataPoints: []metricdata.DataPoint[int64]{
								{
									Value:      val,
									Attributes: *attribute.EmptySet(),
									Time:       time.Now(),
								},
							},
						},
					},
				},
			},
		},
	})
}
