 package response

type EventResponse struct {
	ID          string  `json:"id" bson:"_id"`          // MongoDB ObjectID or PostgreSQL UUID
	Title       string  `json:"title"`                  // Event title
	Description string  `json:"description"`            // Event description
	Price       float64 `json:"price"`                  // Event price
	Location    string  `json:"location"`               // Event location
}

type IdResponse struct {
	ID string `json:"id"` // Unique identifier response (e.g., MongoDB ObjectID)
}
