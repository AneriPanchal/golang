package models

import (
	"time"

	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}
