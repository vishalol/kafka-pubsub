package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

var client *kgo.Client

var topic = "foo"
var group = "my-group-identifier"
var seeds = []string{"localhost:29092", "localhost:19092", "localhost:39092"}

func InitProducer() {
	var err error

	client, err = kgo.NewClient(
		kgo.SeedBrokers(seeds...),
	)
	if err != nil {
		panic(err)
	}
}

func CloseProducer() {
	client.Close()
}

func ProduceMessage(value []byte, key string) error {
	record := kgo.Record{
		Topic: topic,
		Value: value,
		Key:   []byte(key),
	}

	if err := client.ProduceSync(context.Background(), &record).FirstErr(); err != nil {
		return fmt.Errorf("record had a produce error while synchronously producing: %v\n", err)
	}

	return nil
}
