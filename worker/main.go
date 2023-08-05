package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/vishalol/kafka-pubsub/kafka"
	"github.com/vishalol/kafka-pubsub/protomessage"
	"github.com/vishalol/kafka-pubsub/store"
	"google.golang.org/protobuf/proto"
)

func main() {
	numPartitions := 2
	dataStore := store.NewDataStore()

	clients := make([]*kgo.Client, numPartitions)
	for i := 0; i < numPartitions; i++ {
		cl, err := kafka.NewConsumer()
		if err != nil {
			panic(err)
		}
		clients[i] = cl
		go consume(cl, dataStore)
	}

	log.Println("Started consuming messages")

	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, os.Interrupt)

	<-sigs

	log.Println("received interrupt signal; closing client")
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, cl := range clients {
			cl.Close()
		}
	}()

	select {
	case <-sigs:
		log.Println("received second interrupt signal; quitting without waiting for graceful close")
	case <-done:
	}
}

func consume(cl *kgo.Client, dataStore *store.DataStore) {
	for {
		fetches := cl.PollFetches(context.Background())
		if fetches.IsClientClosed() {
			return
		}
		fetches.EachError(func(t string, p int32, err error) {
			die("fetch err topic %s partition %d: %v", t, p, err)
		})

		var seen int
		fetches.EachRecord(func(record *kgo.Record) {
			msg := protomessage.Message{}
			if err := proto.Unmarshal(record.Value, &msg); err != nil {
				log.Println("error while unmarshalling message")
				return
			}

			log.Printf("%s (p=%d): (o=%d)\n", string(record.Key), record.Partition, record.Offset)

			dataStore.Add(msg.Symbol, store.DataPoint{
				Value:     msg.Val,
				Timestamp: msg.Timestamp,
			})

			seen++
		})
		if err := cl.CommitUncommittedOffsets(context.Background()); err != nil {
			log.Printf("commit records failed: %v", err)
			continue
		}
		log.Printf("committed %d records successfully--the recommended pattern, as followed in this demo, is to commit all uncommitted offsets after each poll!\n", seen)
	}
}

func die(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(1)
}
