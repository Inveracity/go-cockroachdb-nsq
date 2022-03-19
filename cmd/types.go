package main

import "github.com/google/uuid"

type task struct {
	ID      uuid.UUID
	JobID   uuid.UUID
	Version string
	Os      string
}
