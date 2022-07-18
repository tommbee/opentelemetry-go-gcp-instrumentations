package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/tommbee/opentelemetry-go-gcp-instrumentations/otelpubsub/example"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topicName := os.Getenv("PUBSUB_TOPIC")
	t, err := example.GetTopic(client, topicName)
	if err != nil {
		log.Fatal(err)
	}
	msg := "Hello World"
	attrs := make(map[string]string)
	attrs["traceparent"] = "00-6ef0bcb5955b1fe1988191a2272577b1-fd61f692a0fdb179-01"
	result := t.Publish(ctx, &pubsub.Message{
		Data:       []byte(msg),
		Attributes: attrs,
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published a message; msg ID: %v\n", id)
}
