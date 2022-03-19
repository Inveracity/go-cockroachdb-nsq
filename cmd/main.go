package main

import (
	"fmt"
	"os"

	"github.com/inveracity/go-cockroachdb-nsq/internal/distributor"
	"github.com/inveracity/go-cockroachdb-nsq/internal/worker"
)

func main() {
	args := os.Args[1:]

	switch arg := args[0]; arg {
	case "distributor":
		distributor.Distributor()
	case "worker":
		worker.Worker()
	default:
		fmt.Printf("one of: distributor, worker, not '%s'\n", arg)
	}
}
