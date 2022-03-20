package task

import (
	"github.com/google/uuid"
)

type Task struct {
	ID      uuid.UUID
	Version string
	Os      string
}
