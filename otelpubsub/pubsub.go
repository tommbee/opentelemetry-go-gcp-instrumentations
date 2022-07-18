package otelpubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "github.com/tommbee/otelpubsub"

type Subscription struct {
	*pubsub.Subscription
	*config
}

// SubscriptionWithTracing creates a new instrumented pubsub.Subscription.
func SubscriptionWithTracing(sub *pubsub.Subscription, opts ...Option) Subscription {
	cfg := config{}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	if cfg.tracerProvider == nil {
		cfg.tracerProvider = otel.GetTracerProvider()
	}
	if cfg.propagators == nil {
		cfg.propagators = otel.GetTextMapPropagator()
	}
	return Subscription{sub, &cfg}
}

// Receive starts a new span before executing the given callback function.
func (t *Subscription) Receive(ctx context.Context, f func(ctx context.Context, m *pubsub.Message)) error {
	subscriptionName := t.Subscription.String()
	tracer := t.config.tracerProvider.Tracer(
		tracerName,
		trace.WithInstrumentationVersion(SemVersion()),
	)

	instrumented := func(msgCtx context.Context, m *pubsub.Message) {
		ctx := PubSubMessageExtractContext(msgCtx, t.config.propagators, m)
		attributes := []attribute.KeyValue{
			semconv.CloudProviderGCP,
			semconv.MessagingOperationReceive,
			semconv.MessagingDestinationKindTopic,
			semconv.MessagingMessageIDKey.String(m.ID),
			semconv.MessagingDestinationKey.String(subscriptionName),
			semconv.MessagingProtocolKey.String("pubsub"),
		}
		opts := []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindConsumer),
			trace.WithAttributes(attributes...),
		}
		ctx, span := tracer.Start(ctx, subscriptionName+" process", opts...)
		defer span.End()
		f(ctx, m)
	}
	return t.Subscription.Receive(ctx, instrumented)
}
