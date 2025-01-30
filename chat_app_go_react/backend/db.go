package routes

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance initializes a connection to the MongoDB database and returns a MongoDB client.
func DBinstance() *mongo.Client {
    MongoDb := "mongodb://localhost:27017"

    client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB")
    return client
}

// Client is a global MongoDB client instance
var Client *mongo.Client = DBinstance()

// OpenConnection creates and returns a collection from the database
func OpenConnection(collectionName string) *mongo.Collection {
    // Use the global Client variable
    collection := Client.Database("caloriesdb").Collection(collectionName)
    return collection
}