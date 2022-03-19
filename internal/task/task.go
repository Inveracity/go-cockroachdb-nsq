package task

import (
	"github.com/google/uuid"
	"github.com/inveracity/go-cockroachdb-nsq/internal/os"
)

type Task struct {
	ID      uuid.UUID
	Version string
	Os      os.OperatingSystem
}
