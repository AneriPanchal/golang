package response

import (
    "eventapp/models"
)

// EventResponse is the response structure for an Event
type EventResponse struct {
	ID          string `json:"id" gorm:"column:id;type:uuid;default:gen_random_uuid()"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"` // Use string to format time for better JSON serialization
}

// FromModel converts an Event model to an EventResponse
func FromModel(event models.Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Date:        event.Date.Format("2006-01-02T15:04:05Z"), // ISO 8601 format
	}
}
