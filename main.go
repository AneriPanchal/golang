package main

import (
	"fmt"
	"log"
	"os"

	"eventapp/config"
	"eventapp/controllers"
	"eventapp/db"
	"eventapp/managers"
	"eventapp/routes"
	"eventapp/service"

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

	// Load PostgreSQL configuration
	postgresCfg, err := config.LoadPostgresConfig()
	if err != nil {
		log.Fatalf("Failed to load PostgreSQL configuration: %v", err)
	}

	// Connect to PostgreSQL
	postgresDB, err := db.ConnectPostgres(postgresCfg)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	// Set PostgreSQL database in the service layer
	service.SetPostgresDB(postgresDB)

	// Connect to MongoDB
	mongoClient, err := db.ConnectMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	// Set MongoDB collection in the service layer
	service.SetEventCollection(mongoClient, "eventapp")

	// Initialize Echo
	e := echo.New()

	// Set the custom validator
	e.Validator = &CustomValidator{Validator: validator.New()}

	// Initialize the manager and controller
	eventManager := &managers.EventManager{}
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
