package service

import (
	"eventapp/models"
	"log"

	"gorm.io/gorm"
)

var postgresDB *gorm.DB

// SetPostgresDB sets the PostgreSQL database connection.
func SetPostgresDB(db *gorm.DB) {
	postgresDB = db
}

// GetEventsFromPostgreSQL retrieves all events from PostgreSQL.
func GetEventsFromPostgreSQL() ([]models.Event, error) {
	var events []models.Event
	err := postgresDB.Find(&events).Error
	if err != nil {
		log.Println("Error fetching events:", err)
		return nil, err
	}
	return events, nil
}

// GetEventByIDFromPostgreSQL retrieves an event by ID from PostgreSQL.
func GetEventByIDFromPostgreSQL(id string) (models.Event, error) {
	var event models.Event
	err := postgresDB.Where("id = ?", id).First(&event).Error
	if err != nil {
		log.Println("Error fetching event by ID:", err)
		return models.Event{}, err
	}
	return event, nil
}

// CreateEventInPostgreSQL creates a new event in PostgreSQL.
func CreateEventInPostgreSQL(event models.Event) (models.Event, error) {
	err := postgresDB.Create(&event).Error
	if err != nil {
		log.Println("Error creating event:", err)
		return models.Event{}, err
	}
	return event, nil
}

// UpdateEventInPostgreSQL updates an existing event in PostgreSQL.
func UpdateEventInPostgreSQL(id string, event models.Event) (models.Event, error) {
	err := postgresDB.Model(&models.Event{}).Where("id = ?", id).Updates(event).Error
	if err != nil {
		log.Println("Error updating event:", err)
		return models.Event{}, err
	}
	event.ID = id
	return event, nil
}

// DeleteEventFromPostgreSQL deletes an event by ID from PostgreSQL.
func DeleteEventFromPostgreSQL(id string) error {
	err := postgresDB.Delete(&models.Event{}, "id = ?", id).Error
	return err
}
