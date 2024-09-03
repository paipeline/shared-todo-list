package models

import (
	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}

var Tasks []Task
