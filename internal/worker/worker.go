package worker

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/inveracity/go-cockroachdb-nsq/internal/task"
	"github.com/nsqio/go-nsq"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	//Process the Message
	var task task.Task
	if err := json.Unmarshal(m.Body, &task); err != nil {
		log.Println("Error when Unmarshaling the message body, Err : ", err)
		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return err
	}

	log.Printf("Received Task\ntask id: %s\nVersion: %s\nOS: %s", task.ID, task.Version, task.Os)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

func Worker() {
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic", "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(&myMessageHandler{})

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
