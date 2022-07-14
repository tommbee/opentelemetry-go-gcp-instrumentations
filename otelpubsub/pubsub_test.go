package otelpubsub

import (
	"cloud.google.com/go/pubsub"
	"testing"
)

func TestSubscriptionWithTracing(t *testing.T) {
	actual := SubscriptionWithTracing(&pubsub.Subscription{})
	if actual.config == nil {
		t.Error("Config was not set")
	}
}
