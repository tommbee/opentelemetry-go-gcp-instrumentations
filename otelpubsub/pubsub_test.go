package otelpubsub

import (
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"context"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestSubscriptionWithTracing(t *testing.T) {
	tracer := trace.NewNoopTracerProvider()
	actual := SubscriptionWithTracing(&pubsub.Subscription{}, WithTracerProvider(tracer))
	assert.Equal(t, tracer, actual.config.tracerProvider)
}

func TestSubscription_Receive(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	otel.SetTracerProvider(sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder)))

	ctx := context.Background()
	srv := pstest.NewServer()
	client, err := pubsub.NewClient(ctx, "test-project",
		option.WithEndpoint(srv.Addr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))
	assert.Nil(t, err)
	defer client.Close()
	defer srv.Close()

	top, err := client.CreateTopic(ctx, "test-topic")
	assert.Nil(t, err)

	pubSubSubConfig := pubsub.SubscriptionConfig{
		Topic:                 top,
		EnableMessageOrdering: true,
	}
	pubSubSub, err := client.CreateSubscription(ctx, "test-subscription", pubSubSubConfig)
	assert.Nil(t, err)

	msg := []byte("Hello")
	publishAttrs := make(map[string]string)
	publishAttrs["traceparent"] = "00-6ef0bcb5955b1fe1988191a2272577b1-fd61f692a0fdb179-01"
	srv.Publish("projects/test-project/topics/test-topic", msg, publishAttrs)
	pubSubSub.ReceiveSettings.Synchronous = true
	propagators := propagation.TraceContext{}
	actual := SubscriptionWithTracing(pubSubSub, WithPropagators(propagators))
	cctx, cancel := context.WithCancel(ctx)
	err = actual.Receive(cctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %q\n", string(m.Data))
		m.Ack()
		cancel()
	})

	if err != nil {
		t.Fatal(err)
	}

	spans := recorder.Ended()
	assert.Len(t, spans, 1)

	span := spans[0]
	parent := span.Parent()
	assert.Equal(t, "projects/test-project/subscriptions/test-subscription process", span.Name())
	assert.Equal(t, trace.SpanKindConsumer, span.SpanKind())
	assert.Equal(t, "6ef0bcb5955b1fe1988191a2272577b1", parent.TraceID().String())
	attrs := span.Attributes()
	assert.Contains(t, attrs, attribute.String("cloud.provider", "gcp"))
	assert.Contains(t, attrs, attribute.String("messaging.operation", "receive"))
	assert.Contains(t, attrs, attribute.String("messaging.destination_kind", "topic"))
	assert.Contains(t, attrs, attribute.String("messaging.protocol", "pubsub"))
}
