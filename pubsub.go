package opentelpubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"go.opentelemetry.io/otel"
)

type Subscription struct {
	pubsub.Subscription
}

type PubSubCarrier struct {
	message *pubsub.Message
}

// Receive calls the configured subscription to trigger the callback on message received.
func (t *Subscription) Receive(ctx context.Context, f func(context.Context, *pubsub.Message)) {
	t.Receive(ctx, f)
}

// NewPubsubCarrier creates a new PubSubCarrier.
func NewPubsubCarrier(msg *pubsub.Message) PubSubCarrier {
	return PubSubCarrier{message: msg}
}

// Get returns the value for a given key
func (c PubSubCarrier) Get(key string) string {
	return c.message.Attributes[key]
}

// Set sets an attribute.
func (c PubSubCarrier) Set(key, val string) {
	c.message.Attributes[key] = val
}

// Keys returns a slice of all keys in the carrier.
func (c PubSubCarrier) Keys() []string {
	i := 0
	out := make([]string, len(c.message.Attributes))
	for k := range c.message.Attributes {
		out[i] = k
		i++
	}
	return out
}

func PubSubMessageInjectContext(ctx context.Context, msg *pubsub.Message) {
	otel.GetTextMapPropagator().Inject(ctx, NewPubsubCarrier(msg))
}

func PubSubMessageExtractContext(ctx context.Context, msg *pubsub.Message) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, NewPubsubCarrier(msg))
}
