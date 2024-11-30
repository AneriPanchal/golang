package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          string    `bson:"_id,omitempty" json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Title       string    `bson:"title" json:"title" gorm:"column:title" validate:"required"`
	Description string    `bson:"description" json:"description" gorm:"column:description" validate:"required"`
	Date        time.Time `bson:"date" json:"date" gorm:"column:date" validate:"required"`
}

// SetMongoID assigns a MongoDB ObjectID to the Event if it's not already set.
func (e *Event) SetMongoID() {
	if e.ID == "" {
		e.ID = primitive.NewObjectID().Hex()
	}
}

// GenerateUUID assigns a UUID to the Event if it's not already set (for PostgreSQL).
func (e *Event) GenerateUUID() {
	if e.ID == "" {
		e.ID = uuid.New().String()
	}
}
