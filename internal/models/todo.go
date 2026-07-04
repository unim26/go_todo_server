package models

import "github.com/google/uuid"

type TODO struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
}
