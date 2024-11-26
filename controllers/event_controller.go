package controllers

import (
	"net/http"
	"eventapp/managers"
	"eventapp/models"

	"github.com/labstack/echo/v4"
)

// EventController handles HTTP requests for events
type EventController struct {
	Manager *managers.EventManager
}

// GetEvents fetches all events
func (c *EventController) GetEvents(ctx echo.Context) error {
	// Parse the flag from query parameters (true for MongoDB, false for PostgreSQL)
	flag := ctx.QueryParam("useMongo") == "true"

	// Call the manager to get events
	events, err := c.Manager.GetEvents(flag)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, events)
}

// GetEventByID fetches a single event by its ID
func (c *EventController) GetEventByID(ctx echo.Context) error {
	id := ctx.Param("id")
	flag := ctx.QueryParam("useMongo") == "true"

	event, err := c.Manager.GetEventByID(flag, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, event)
}

// CreateEvent creates a new event
func (c *EventController) CreateEvent(ctx echo.Context) error {
	var request models.Event
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Use the flag to determine the database
	flag := ctx.QueryParam("useMongo") == "true"

	// FIX: Pass both the flag and the request to CreateEvent
	event, err := c.Manager.CreateEvent(flag, request) // Corrected to include the flag
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, event)
}

// UpdateEvent updates an existing event
func (c *EventController) UpdateEvent(ctx echo.Context) error {
	id := ctx.Param("id")
	var request models.Event
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Use the flag to determine the database
	flag := ctx.QueryParam("useMongo") == "true"

	// Call the manager to update the event
	event, err := c.Manager.UpdateEvent(flag, id, request) // Corrected to include the flag
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, event)
}

// DeleteEvent deletes an event by its ID
func (c *EventController) DeleteEvent(ctx echo.Context) error {
	id := ctx.Param("id")

	// Use the flag to determine the database
	flag := ctx.QueryParam("useMongo") == "true"

	// Call the manager to delete the event
	if err := c.Manager.DeleteEvent(flag, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
