package mongoConnector

import (
	"context"
	"fmt"
	"log"
	"time"

	"../dtos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context

func Insert(learning dtos.Learning) bool {
	openConnection()
	log.Println(learning)
	closeConnection()

	return true
}

func Search(text string) dtos.Learning {
	openConnection()
	log.Println(text)
	closeConnection()

	return dtos.Learning{Intent: text}
}

func openConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func closeConnection() {
	if client == nil {
		log.Println("nulo")
	}
	err := client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
