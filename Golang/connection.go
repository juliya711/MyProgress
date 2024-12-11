package main

import (
    "context" // Import the context package
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Get the MongoDB URI and Database Name from environment variables
    mongoURI := os.Getenv("MONGO_URI")
    databaseName := os.Getenv("DATABASE_NAME")

    // Connect to MongoDB
    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions) // Use context.TODO()
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil) // Use context.TODO()
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    fmt.Println("Connected to MongoDB!")
    fmt.Printf("Using database: %s\n", databaseName)
}