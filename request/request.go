// package request

// // CreateRequest defines the structure for the event creation request
// type CreateRequest struct {
// 	Title       string  `json:"title" validate:"required"`        // Event title
// 	Date        string  `json:"date" validate:"required"`         // Event date
// 	Location    string  `json:"location" validate:"required"`     // Event location
// 	Description string  `json:"description" validate:"required"`  // Event description
// 	Price       float64 `json:"price" validate:"required,gt=0"`    // Event price (optional, must be greater than 0)
// }

// // UpdateRequest defines the structure for the event update request
// type UpdateRequest struct {
// 	Title       string  `json:"title"`       // Event title (optional)
// 	Date        string  `json:"date"`        // Event date (optional)
// 	Location    string  `json:"location"`    // Event location (optional)
// 	Description string  `json:"description"` // Event description (optional)
// 	Price       float64 `json:"price"`       // Event price (optional)
// }

// // DeleteRequest defines the structure for the event deletion request
// type DeleteRequest struct {
// 	ID int `json:"id" validate:"required"` // Event ID for deletion
// }
package request

import (
	"github.com/go-playground/validator/v10"
	//"fmt"
)

var validate = validator.New()

type CreateEventRequest struct {
	ID          string  `json:"id" bson:"_id"`
	Title       string  `json:"title" validate:"required,min=1,max=100"`  // Event title
	Description string  `json:"description" validate:"required"`          // Event description
	Price       float64 `json:"price" validate:"required,gt=0"`           // Event price
	Location    string  `json:"location" validate:"required,min=1,max=255"` // Event location
}

type UpdateEventRequest struct {
	ID          string  `json:"id" bson:"_id"`
	Title       string  `json:"title" validate:"required,min=1,max=100"`  // Event title
	Description string  `json:"description" validate:"required"`          // Event description
	Price       float64 `json:"price" validate:"required,gt=0"`           // Event price
	Location    string  `json:"location" validate:"required,min=1,max=255"` // Event location
}

//Validate validates the request payload.
func (r *CreateEventRequest) Validate() error {
	return validate.Struct(r)
}

func (r *UpdateEventRequest) Validate() error {
	return validate.Struct(r)
}