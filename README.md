# Go Open Telemetry Instrumentation for GCP

A collection of 3rd party packages for use with GCP & [Open Telemetry Go](https://github.com/open-telemetry/opentelemetry-go)

## PubSub

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