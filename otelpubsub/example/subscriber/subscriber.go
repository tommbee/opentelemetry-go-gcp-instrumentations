package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/tommbee/opentelemetry-go-gcp-instrumentations/otelpubsub"
	"github.com/tommbee/opentelemetry-go-gcp-instrumentations/otelpubsub/example"
	"log"
	"os"
)

func main() {
	example.InitTracer()
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := os.Getenv("PUBSUB_TOPIC")
	topic, err := example.GetTopic(client, topicName)
	if err != nil {
		log.Println(err)
	}

	subName := os.Getenv("PUBSUB_SUBSCRIPTION_NAME")
	pubSubSub, ex := example.GetSubscription(client, subName, topic)
	if ex != nil {
		log.Println(err)
	}

	sub := otelpubsub.NewSubscriptionWithTracing(pubSubSub)
	e := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %q\n", string(m.Data))
		m.Ack()
	})
	if e != nil {
		log.Println(err)
	}
}
