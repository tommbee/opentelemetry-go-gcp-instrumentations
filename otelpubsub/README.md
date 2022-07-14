# PubSub Instrumentation

## Subscriptions

Trace messages received by a subscription using the implementation below.

```go
package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/tommbee/opentelemetry-go-gcp-instrumentations/otelpubsub"
	"log"
)

func main() {
	InitTracer()
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "example-project")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
	pubSubSub := client.Subscription("example-subscription")
	sub := otelpubsub.NewSubscriptionWithTracing(pubSubSub)
	
	e := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %q\n", string(m.Data))
		m.Ack()
	})
	if e != nil {
		log.Println(err)
	}
}
```
