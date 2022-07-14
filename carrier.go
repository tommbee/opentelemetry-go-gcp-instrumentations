package otelpubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"go.opentelemetry.io/otel"
)

type PubSubMessageCarrier struct {
	message *pubsub.Message
}

// NewPubSubMessageCarrier creates a new PubSubMessageCarrier.
func NewPubSubMessageCarrier(msg *pubsub.Message) PubSubMessageCarrier {
	return PubSubMessageCarrier{message: msg}
}

// Get returns the value for a given key.
func (c PubSubMessageCarrier) Get(key string) string {
	return c.message.Attributes[key]
}

// Set the value for a given key.
func (c PubSubMessageCarrier) Set(key, val string) {
	c.message.Attributes[key] = val
}

// Keys returns a slice of all keys in the carrier.
func (c PubSubMessageCarrier) Keys() []string {
	i := 0
	out := make([]string, len(c.message.Attributes))
	for k := range c.message.Attributes {
		out[i] = k
		i++
	}
	return out
}

// PubSubMessageExtractContext sets NewPubSubMessageCarrier as an "extract" concern for the global propagator
func PubSubMessageExtractContext(ctx context.Context, msg *pubsub.Message) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, NewPubSubMessageCarrier(msg))
}
