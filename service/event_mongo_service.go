package service

import (
	"context"
	"eventapp/models"
	"log"
	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var eventCollection *mongo.Collection

// SetEventCollection sets the MongoDB collection for events.
func SetEventCollection(client *mongo.Client, database string) {
	eventCollection = client.Database(database).Collection("events")
}

// GetEventsFromMongoDB retrieves all events from MongoDB.
func GetEventsFromMongoDB() ([]models.Event, error) {
	cursor, err := eventCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("Error fetching events:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var events []models.Event
	for cursor.Next(context.Background()) {
		var event models.Event
		if err := cursor.Decode(&event); err != nil {
			log.Println("Error decoding event:", err)
			continue
		}
		events = append(events, event)
	}
	return events, nil
}

// GetEventByIDFromMongoDB retrieves an event by ID from MongoDB.
func GetEventByIDFromMongoDB(id primitive.ObjectID) (models.Event, error) {
	var event models.Event
	err := eventCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		log.Println("Error fetching event by ID:", err)
		return models.Event{}, err
	}
	return event, nil
}

// CreateEventInMongoDB creates a new event in MongoDB.
func CreateEventInMongoDB(event models.Event) (models.Event, error) {
	result, err := eventCollection.InsertOne(context.Background(), event)
	if err != nil {
		return models.Event{}, err
	}
	event.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return event, nil
}

// UpdateEventInMongoDB updates an existing event in MongoDB.
func UpdateEventInMongoDB(id primitive.ObjectID, event models.Event) (models.Event, error) {
	_, err := eventCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": event},
	)
	if err != nil {
		return models.Event{}, err
	}
	event.ID = id.Hex()
	return event, nil
}

// DeleteEventFromMongoDB deletes an event by ID from MongoDB.
func DeleteEventFromMongoDB(id primitive.ObjectID) error {
	_, err := eventCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
