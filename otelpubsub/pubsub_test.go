package otelpubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log"
	"testing"
)

func TestSubscriptionWithTracing(t *testing.T) {
	tracer := trace.NewNoopTracerProvider()
	propagators := otel.GetTextMapPropagator()
	actual := SubscriptionWithTracing(&pubsub.Subscription{}, WithTracerProvider(tracer), WithPropagators(propagators))

	if actual.config.tracerProvider != tracer {
		t.Error("Trace provider not set")
	}

	if actual.config.propagators != propagators {
		t.Error("Propagators not set")
	}
}

func TestSubscription_Receive(t *testing.T) {
	actual := SubscriptionWithTracing(&pubsub.Subscription{})
	err := actual.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %q\n", string(m.Data))
		m.Ack()
	})

	if err != nil {
		t.Fatal(err)
	}
}
