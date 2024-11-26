package main

import (
	"fmt"
	"log"
	"os"
	"eventapp/config"
	"eventapp/controllers"
	"eventapp/managers"
	"eventapp/routes"
	"eventapp/service"
	"eventapp/models"

	_ "github.com/lib/pq"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// CustomValidator implements Echo's Validator interface
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate validates the struct using go-playground/validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configurations.")
	}

	// Initialize database connections
	if err := config.InitializeDatabaseConnections(); err != nil {
		log.Fatalf("Failed to initialize database connections: %v", err)
	}

	// Set up MongoDB and PostgreSQL services
	service.SetEventCollection(config.MongoClient, "eventapp")
	service.SetPostgresDB(config.PostgresDB)

	// Initialize Echo
	e := echo.New()

	// Set the custom validator
	e.Validator = &CustomValidator{Validator: validator.New()}

	// Initialize the manager
	eventManager := &managers.EventManager{}

	// Initialize the controller
	eventController := &controllers.EventController{Manager: eventManager}

	// Set up routes
	routes.SetupRoutes(e, eventController)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	address := fmt.Sprintf(":%s", port)

	log.Printf("Starting server on %s", address)
	if err := e.Start(address); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
