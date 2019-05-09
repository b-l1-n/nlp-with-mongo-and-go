package mongoConnector

import (
	"context"
	"fmt"
	"log"
	"time"
	"agent/dtos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Insert(learning dtos.Learning) bool {
	client, ctx := openConnection()
	log.Println(learning)
	closeConnection(client, ctx)

	return true
}

func Search(text string) dtos.Learning {
	client, ctx := openConnection()
	log.Println(text)
	closeConnection(client, ctx)

	return dtos.Learning{Intent: text}
}

func openConnection() (*mongo.Client, context.Context) {
	mongoHost, exists := os.LookupEnv("MONGO_HOST")
    if !exists {
        mongoHost = "localhost"
    }

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + mongoHost + ":27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, ctx
}

func closeConnection(client *mongo.Client, ctx context.Context) {
	err := client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
