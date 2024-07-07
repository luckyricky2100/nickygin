package logger

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func NewGormLog(ctx context.Context) trace.Span {
	tracer := otel.Tracer("gorm.io/plugin/opentelemetry")

	ctx, span := tracer.Start(ctx, "root")
	return span
}

func PrintTraceID(ctx context.Context) {
	fmt.Println("trace:", TraceURL(trace.SpanFromContext(ctx)))
}
func TraceURL(span trace.Span) string {
	switch {
	case os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT") != "":
		return fmt.Sprintf("http://localhost:16686/trace/%s", span.SpanContext().TraceID())
	default:
		return fmt.Sprintf("http://localhost:16686/trace/%s", span.SpanContext().TraceID())
	}
}
