package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	nsq "github.com/nsqio/go-nsq"

	os "github.com/inveracity/go-cockroachdb-nsq/internal"
)

func distributor() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	msg := Task{
		ID:      uuid.Must(uuid.NewRandom()),
		Version: "2022.1",
		Os:      os.Linux,
	}
	//Convert message as []byte
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	topicName := "topic"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, payload)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()
}
