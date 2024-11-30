package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"eventapp/config"
	"eventapp/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Client *mongo.Client
	PG     *gorm.DB
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// ConnectPostgres establishes a connection to PostgreSQL using configuration.
func ConnectPostgres(cfg *config.PostgresConfig) (*gorm.DB, error) {
	postgresDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.POSTGRES_HOST,
		cfg.POSTGRES_PORT,
		cfg.POSTGRES_USER,
		cfg.POSTGRES_PASSWORD,
		cfg.POSTGRES_DB,
	)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully!")

	// Perform automatic migrations for the Event model
	if err := db.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("Error during auto-migration: %v", err)
		return nil, err
	}
	log.Println("Database migration completed!")

	PG = db
	return PG, nil
}

// ConnectMongoDB initializes the MongoDB client.
func ConnectMongoDB() (*mongo.Client, error) {
	LoadEnv()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Ping the MongoDB database to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB successfully!")
	Client = client
	return client, nil
}
