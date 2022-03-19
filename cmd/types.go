package main

import (
	"github.com/google/uuid"
	os "github.com/inveracity/go-cockroachdb-nsq/internal"
)

type Task struct {
	ID      uuid.UUID
	Version string
	Os      os.OperatingSystem
}
