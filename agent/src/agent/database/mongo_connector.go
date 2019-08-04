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
var entitiesCollection string = "entities"

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

func InsertEntity(entity dtos.Entity) bool {
	client, ctx := openConnection()
	var operationStatus bool = true
	log.Println("Trying to insert Learning object into Database")
	collection := client.Database(intentDatabase).Collection(entitiesCollection)

	if collection == nil {
		log.Println("No collection with name " + entitiesCollection + " in database " + intentDatabase)
		operationStatus = false
	}

	for _, entityValue := range entity.Values {
		entityDB := dtos.EntityDB{}
		entityDB.Name = entity.Name
		entityDB.Values = entityValue

		_, err := collection.InsertOne(ctx, entityDB)

		if err != nil {
			operationStatus = false
			log.Println(err)
		}

		if operationStatus {
			log.Println("Learning object inserted into Database")
		}
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

func RetrieveEntities() []*dtos.Entity {
	
	client, ctx := openConnection()
	log.Println("Trying to retrieve Entities")
	collection := client.Database(intentDatabase).Collection(entitiesCollection)

	if collection == nil {
		log.Println("No collection with name " + entitiesCollection + " in database " + intentDatabase)
		return nil
	}

	var entities []*dtos.Entity
	filter := bson.D{{ }}
	cur, err := collection.Find(ctx, filter)

	// Loop over the cursor retrieved
	for cur.Next(ctx) {
    
		// create a value into which the single document can be decoded
		var entityDB dtos.EntityDB
		err := cur.Decode(&entityDB)
		if err != nil {
			log.Fatal(err)
		}
		
		isNew := true;
		for _, externalEntity := range entities {
			if externalEntity.Name == entityDB.Name {
				externalEntity.Values = append(externalEntity.Values, entityDB.Values)
				isNew = false
			}
		}

		if isNew {
			var externalEntity dtos.Entity
			externalEntity.Name = entityDB.Name
			externalEntity.Values = []string{entityDB.Values}
			entities = append(entities, &externalEntity)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	
	// Close the cursor once finished
	cur.Close(ctx)

	if err != nil {
		log.Println(err)
		return nil
	}
	closeConnection(client, ctx)

	return entities
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

// db.entities.find( {$text : { $search : "Me gusta Die Hard los lunes" } }, {_id: 0 , Name : 1 ,Values : { $elemMatch : { $regex : /(Me gusta Die Hard los lunes).toLowerCase()/i }}})

func ExtractEntities(text string) []*dtos.EntityDB {
	client, ctx := openConnection()
	log.Println("Trying to extract Entities from utterance : " + text)
	collection := client.Database(intentDatabase).Collection(entitiesCollection)

	if collection == nil {
		log.Println("No collection with name " + entitiesCollection + " in database " + intentDatabase)
		return nil
	}

	var result []*dtos.EntityDB
	filter := bson.D{{ "$text" , bson.D{{ "$search" , text }}  }}
	cur, err := collection.Find(ctx, filter)

	for cur.Next(ctx) {
		var entityDB dtos.EntityDB
		err := cur.Decode(&entityDB)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, &entityDB)
	}
		
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	
	cur.Close(ctx)

	if err != nil {
		log.Println(err)
		return nil
	}

	closeConnection(client, ctx)

	return result
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
