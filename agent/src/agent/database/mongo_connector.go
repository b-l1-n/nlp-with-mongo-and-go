package mongoConnector

import (
	"agent/dtos"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var intentDatabase string = "intents_db"
var intentsCollection string = "intents"

func InsertLearning(learning dtos.Learning) bool {
	client, ctx := openConnection()
	var operationStatus bool = true
	log.Println("Trying to insert Learning object into Database")
	collection := client.Database(intentDatabase).Collection(intentsCollection)

	if collection == nil {
		log.Println("No collection with name " + intentsCollection + " in database " + intentDatabase)
		operationStatus = false
	}

	_, err := collection.InsertOne(ctx, learning)

	if err != nil {
		operationStatus = false
		log.Println(err)
	}

	if operationStatus {
		log.Println("Learning object inserted into Database")
	}

	closeConnection(client, ctx)

	return operationStatus
}

func RetrieveLearning(intentName string) *dtos.Learning {
	

	client, ctx := openConnection()
	log.Println("Trying to retrieve Learning with intent name: " + intentName)
	collection := client.Database(intentDatabase).Collection(intentsCollection)

	if collection == nil {
		log.Println("No collection with name " + intentsCollection + " in database " + intentDatabase)
		return nil
	}

	var result dtos.Learning
	filter := bson.D{{ "Intent" , bson.D{{ "$eq" , intentName }}  }}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil
	}
	closeConnection(client, ctx)

	return &result
}

func Search(text string) *dtos.Learning {
	client, ctx := openConnection()
	log.Println("Trying to retrieve Response to utterance : " + text)
	collection := client.Database(intentDatabase).Collection(intentsCollection)

	if collection == nil {
		log.Println("No collection with name " + intentsCollection + " in database " + intentDatabase)
		return nil
	}

	var result dtos.Learning
	filter := bson.D{{ "$text" , bson.D{{ "$search" , text }}  }}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil
	}
	closeConnection(client, ctx)

	return &result
}

func openConnection() (*mongo.Client, context.Context) {
	mongoHost, exists := os.LookupEnv("MONGO_HOST")
	if !exists {
		mongoHost = "localhost"
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongoHost+":27017"))

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
