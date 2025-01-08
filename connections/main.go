package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Juliya:jp12345@cluster0.gmcyl.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect client: %v", err)
		}
	}()

	// Ping the primary node
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	// List databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to list databases: %v", err)
	}
	fmt.Println("Databases:", databases)
}
