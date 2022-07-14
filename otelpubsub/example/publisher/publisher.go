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
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Published a message; msg ID: %v\n", id)
}
