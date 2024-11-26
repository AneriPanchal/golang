package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MongoClient *mongo.Client
	PostgresDB  *gorm.DB
)

// LoadConfig loads environment variables from the `.env` file.
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configurations.")
	}
}

// GetEnv retrieves the value of an environment variable.
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// ConnectMongoDB establishes a connection to MongoDB.
func ConnectMongoDB() (*mongo.Client, error) {
	LoadConfig()

	mongoURI := GetEnv("MONGO_URI", "mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(mongoURI).SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("Failed to connect to MongoDB:", err)
		return nil, err
	}

	// Ping the database to verify the connection
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Println("MongoDB ping failed:", err)
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return client, nil
}

// ConnectPostgres establishes a connection to PostgreSQL.
func ConnectPostgres() (*gorm.DB, error) {
	LoadConfig()

	// Fetch database connection details from the environment
	host := GetEnv("POSTGRES_HOST", "localhost")
	port := GetEnv("POSTGRES_PORT", "5433")
	user := GetEnv("POSTGRES_USER", "postgres")
	password := GetEnv("POSTGRES_PASSWORD", "admin")
	dbname := GetEnv("POSTGRES_DB", "eventapp")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a GORM connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to PostgreSQL:", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL")
	return db, nil
}

// InitializeDatabaseConnections initializes both MongoDB and PostgreSQL connections.
func InitializeDatabaseConnections() error {
	var err error

	// Connect to MongoDB
	MongoClient, err = ConnectMongoDB()
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Connect to PostgreSQL
	PostgresDB, err = ConnectPostgres()
	if err != nil {
		return fmt.Errorf("error connecting to PostgreSQL: %w", err)
	}

	return nil
}
