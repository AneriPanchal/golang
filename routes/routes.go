package routes

import (
	"eventapp/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, controller *controllers.EventController) {
	e.GET("/events", controller.GetEvents)
	e.GET("/events/:id", controller.GetEventByID)
	e.POST("/events", controller.CreateEvent)
	e.PUT("/events/:id", controller.UpdateEvent)
	e.DELETE("/events/:id", controller.DeleteEvent)
}
