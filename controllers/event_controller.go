package controllers

import (
	"log"
	"net/http"
	"eventapp/managers"
	"eventapp/request"
	"eventapp/response"

	"github.com/labstack/echo/v4"
)

// EventController handles HTTP requests related to Events
type EventController struct {
	Manager *managers.EventManager
}

// GetEvents handler to fetch all Events
func (c *EventController) GetEvents(ctx echo.Context) error {
	flagValue := ctx.QueryParam("useMongo")
	flag := flagValue == "true"

	events, err := c.Manager.GetEvents(flag)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	eventResponses := make([]response.EventResponse, 0, len(events))
	for _, event := range events {
		eventResponses = append(eventResponses, response.FromModel(event))
	}

	log.Println("Returned all events")
	return ctx.JSON(http.StatusOK, eventResponses)
}

// GetEventByID handler to fetch an Event by ID
func (c *EventController) GetEventByID(ctx echo.Context) error {
	id := ctx.Param("id")
	flagValue := ctx.QueryParam("useMongo")
	flag := flagValue == "true"

	event, err := c.Manager.GetEventByID(flag, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Printf("Returned event with ID %s", id)
	return ctx.JSON(http.StatusOK, response.FromModel(event))
}

// CreateEvent handler to create a new Event
func (c *EventController) CreateEvent(ctx echo.Context) error {
	flagValue := ctx.QueryParam("useMongo")
	flag := flagValue == "true"

	var req request.EventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	event := req.ToModel()
	createdEvent, err := c.Manager.CreateEvent(flag, event)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Println("Created new event")
	return ctx.JSON(http.StatusCreated, response.FromModel(createdEvent))
}

// UpdateEvent handler to update an existing Event
func (c *EventController) UpdateEvent(ctx echo.Context) error {
	id := ctx.Param("id")
	flagValue := ctx.QueryParam("useMongo")
	flag := flagValue == "true"

	var req request.EventRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	event := req.ToModel()
	updatedEvent, err := c.Manager.UpdateEvent(flag, id, event)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Printf("Updated event with ID %s", id)
	return ctx.JSON(http.StatusOK, response.FromModel(updatedEvent))
}

// DeleteEvent handler to delete an Event
func (c *EventController) DeleteEvent(ctx echo.Context) error {
	id := ctx.Param("id")
	flagValue := ctx.QueryParam("useMongo")
	flag := flagValue == "true"

	if err := c.Manager.DeleteEvent(flag, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	log.Printf("Deleted event with ID %s", id)
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Event deleted successfully"})
}
