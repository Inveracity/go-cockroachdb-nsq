package distributor

import (
	"context"
	"encoding/json"
	"log"

	crud "github.com/inveracity/go-cockroachdb-nsq/internal/database"
	nsq "github.com/nsqio/go-nsq"
)

func Distributor(ctx context.Context) {

	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := crud.Database()
	conn := crud.NewDB(db)
	res, _ := conn.Read(ctx)
	payload, err := json.Marshal(res)
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
