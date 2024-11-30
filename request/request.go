package request

import (
    "eventapp/models" 
	"time"
)

// EventRequest is the request structure for creating and updating an Event
type EventRequest struct {
	ID          string    `bson:"_id,omitempty" json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Title       string    `bson:"title" json:"title" validate:"required" gorm:"column:title"`
	Description string    `bson:"description" json:"description" validate:"required" gorm:"column:description"`
	Date        time.Time `bson:"date" json:"date" validate:"required" gorm:"column:date"`
}

// ToModel converts an EventRequest to an Event model
func (req *EventRequest) ToModel() models.Event {
	return models.Event{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
	}
}
