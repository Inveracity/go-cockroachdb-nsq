package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	crud "github.com/inveracity/go-cockroachdb-nsq/internal/database"
	"github.com/inveracity/go-cockroachdb-nsq/internal/distributor"
	operatingsystem "github.com/inveracity/go-cockroachdb-nsq/internal/os"
	"github.com/inveracity/go-cockroachdb-nsq/internal/task"
	"github.com/inveracity/go-cockroachdb-nsq/internal/worker"
)

func Init() {

	_, err := crud.CreateDatabase("work")
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
	}

	_, err = crud.CreateTable("work")
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
	}

	db, err := crud.Database()
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
	}
	newdb := crud.NewDB(db)
	newtask := task.Task{
		ID:      uuid.Must(uuid.NewRandom()),
		Version: "2022.1",
		Os:      operatingsystem.Linux,
	}
	err = newdb.Insert(context.Background(), newtask)
	if err != nil {
		log.Printf("ERROR: %+v\n", err)
	}
}

func main() {
	ctx := context.Background()
	args := os.Args[1:]

	switch arg := args[0]; arg {
	case "distributor":
		distributor.Distributor(ctx)
	case "worker":
		worker.Worker()
	case "init":
		Init()
	default:
		fmt.Printf("one of: distributor, worker, not '%s'\n", arg)
	}
}
