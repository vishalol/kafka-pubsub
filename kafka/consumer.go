package kafka

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewConsumer() (*kgo.Client, error) {
	// One client can both produce and consume!
	// Consuming can either be direct (no consumer group), or through a group. Below, we use a group.
	client, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topic),
		kgo.DisableAutoCommit(),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
