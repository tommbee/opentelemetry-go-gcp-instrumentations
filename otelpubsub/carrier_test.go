package otelpubsub

import (
	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPubSubMessageCarrier_Keys(t *testing.T) {
	msg := pubsub.Message{Attributes: map[string]string{
		"foo": "bar",
	}}
	carrier := NewPubSubMessageCarrier(&msg)
	assert.Equal(t, []string([]string{"foo"}), carrier.Keys())
}

func TestPubSubMessageCarrier_Get(t *testing.T) {
	msg := pubsub.Message{Attributes: map[string]string{
		"foo": "bar",
	}}
	carrier := NewPubSubMessageCarrier(&msg)
	assert.Equal(t, "bar", carrier.Get("foo"))
}

func TestPubSubMessageCarrier_Set(t *testing.T) {
	msg := pubsub.Message{Attributes: map[string]string{
		"foo": "bar",
	}}
	carrier := NewPubSubMessageCarrier(&msg)
	carrier.Set("baz", "boo")
	assert.Equal(t, "boo", msg.Attributes["baz"])
}
