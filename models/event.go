package models

import (
	"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title" validate:"required"`
	Description string    `json:"description" bson:"description" validate:"required"`
	Date        time.Time `json:"date" bson:"date" validate:"required"`
}
