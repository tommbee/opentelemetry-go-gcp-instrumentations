package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

func GetTopic(client *pubsub.Client, name string) (*pubsub.Topic, error) {
	log.Print("Creating topic...")
	ctx := context.Background()
	var err error
	topic := client.Topic(name)
	exists, err := topic.Exists(ctx)

	if exists != true {
		topic, err = client.CreateTopic(ctx, name)
	}

	return topic, err
}

func GetSubscription(client *pubsub.Client, name string, topic *pubsub.Topic) (*pubsub.Subscription, error) {
	log.Print("Creating subscription...")
	ctx := context.Background()
	var err error
	sub := client.Subscription(name)
	exists, err := sub.Exists(ctx)

	if exists != true {
		sub, err = client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{Topic: topic})
	}

	return sub, err
}
