package main

import (
	"github.com/inveracity/go-cockroachdb-nsq/internal/distributor"
	"github.com/inveracity/go-cockroachdb-nsq/internal/worker"
)

func main() {
	distributor.Distributor()
	worker.Worker()
}
