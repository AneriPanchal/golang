package managers

import (
	//"errors"
	"eventapp/models"
	"eventapp/request"
	"eventapp/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventManager struct{}

// GetEvents retrieves all events from the specified data source.
func (m *EventManager) GetEvents(flag bool) ([]models.Event, error) {
	if flag {
		return service.GetEventsFromMongoDB()
	}
	return service.GetEventsFromPostgreSQL()
}

// GetEventByID retrieves an event by its ID from the specified data source.
func (m *EventManager) GetEventByID(flag bool, id string) (models.Event, error) {
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return models.Event{}, err
		}
		return service.GetEventByIDFromMongoDB(objectID)
	}
	return service.GetEventByIDFromPostgreSQL(id)
}

// CreateEvent creates a new event in the specified data source.
func (m *EventManager) CreateEvent(flag bool, event request.EventRequest) (models.Event, error) {
	event1 := event.ToModel()

	if flag {
		return service.CreateEventInMongoDB(event1)
	}
	return service.CreateEventInPostgreSQL(event1)
}

// UpdateEvent updates an existing event in the specified data source.
func (m *EventManager) UpdateEvent(flag bool, id string, event request.EventRequest) (models.Event, error) {
	event1 := event.ToModel()
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return models.Event{}, err
		}
		return service.UpdateEventInMongoDB(objectID, event1)
	}
	return service.UpdateEventInPostgreSQL(id, event1)
}

// DeleteEvent deletes an event from the specified data source.
func (m *EventManager) DeleteEvent(flag bool, id string) error {
	if flag {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		return service.DeleteEventFromMongoDB(objectID)
	}
	return service.DeleteEventFromPostgreSQL(id)
}
