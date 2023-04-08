package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type InterceptorsMetricsFactory interface {
	MetricsServerInterceptor(service string) grpc.UnaryServerInterceptor
	MetricsClientInterceptor(client string) grpc.UnaryClientInterceptor
	TracingServerInterceptor() grpc.UnaryServerInterceptor
	TracingClientInterceptor() grpc.UnaryClientInterceptor
	Listen()
}

func New(namespace string, metricsPort string, tracesUrl string, log *zap.Logger) InterceptorsMetricsFactory {
	imp := &implementation{namespace, metricsPort, log}
	log.Info("tracesUrl", zap.String("tracesUrl", tracesUrl))
	imp.initTracer(tracesUrl)
	return imp
}

type implementation struct {
	namespace   string
	metricsPort string
	log         *zap.Logger
}

func (f *implementation) initTracer(url string) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		f.log.Fatal("cannot listen http", zap.Error(err))
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(f.namespace),
			attribute.String("environment", "production"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

}

func (f *implementation) TracingServerInterceptor() grpc.UnaryServerInterceptor {
	return otelgrpc.UnaryServerInterceptor()
}

func (f *implementation) TracingClientInterceptor() grpc.UnaryClientInterceptor {
	return otelgrpc.UnaryClientInterceptor()
}

func (f *implementation) Listen() {
	http.Handle("/metrics", promhttp.Handler())
	f.log.Info("metrics listening http", zap.String("metricsPort", f.metricsPort))
	err := http.ListenAndServe(":"+f.metricsPort, nil)
	f.log.Fatal("cannot listen http", zap.Error(err))
}

func (f *implementation) MetricsServerInterceptor(service string) grpc.UnaryServerInterceptor {
	var (
		RequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: f.namespace,
			Subsystem: service,
			Name:      "requests_total",
		},
			[]string{"method"},
		)
		ResponseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: f.namespace,
			Subsystem: service,
			Name:      "responses_total",
		},
			[]string{"method", "status"},
		)
		HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: f.namespace,
			Subsystem: service,
			Name:      "histogram_response_time_seconds",
			Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
		},
			[]string{"method"},
		)
	)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		RequestsCounter.WithLabelValues(info.FullMethod).Inc()

		timeStart := time.Now()

		res, err := handler(ctx, req)
		elapsed := time.Since(timeStart)

		if err == nil {
			HistogramResponseTime.WithLabelValues(info.FullMethod).Observe(elapsed.Seconds())
			ResponseCounter.WithLabelValues(info.FullMethod, "success").Inc()
		} else {
			ResponseCounter.WithLabelValues(info.FullMethod, "error").Inc()
		}
		return res, err
	}

}

func (f *implementation) MetricsClientInterceptor(service string) grpc.UnaryClientInterceptor {
	var (
		ResponseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: f.namespace,
			Subsystem: service,
			Name:      "responses_total",
		},
			[]string{"method", "status"},
		)
		HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: f.namespace,
			Subsystem: service,
			Name:      "histogram_response_time_seconds",
			Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
		},
			[]string{"method"},
		)
	)

	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		elapsed := time.Since(start)

		if err == nil {
			HistogramResponseTime.WithLabelValues(method).Observe(elapsed.Seconds())
			ResponseCounter.WithLabelValues(method, "success").Inc()
		} else {
			ResponseCounter.WithLabelValues(method, "error").Inc()
		}
		return err
	}
}
