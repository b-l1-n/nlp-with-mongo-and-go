package agentHandler

import (
	"agent/dtos"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"agent/database"
)

func DetectIntent(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		searchInLearningDatabase(response, request)
	default:
		returnMessage(response, request, "Method not allowed")
	}
}

func LearnUtterance(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		insertIntoLearningDatabase(response, request)
	case "GET":
		retrieveLearningFromDatabase(response, request)
	default:
		returnMessage(response, request, "Method not allowed")
	}
}

func insertIntoLearningDatabase(response http.ResponseWriter, request *http.Request) {
	learning := dtos.Learning{}

	error := json.NewDecoder(request.Body).Decode(&learning)

	if error != nil {
		panic(error)
	}

	ok := mongoConnector.InsertLearning(learning)

	if ok {
		conversion, error := json.Marshal(learning)

		if error != nil {
			panic(error)
		}

		response.Write(conversion)
	}

	response.Write(nil)
}

func retrieveLearningFromDatabase(response http.ResponseWriter, request *http.Request) {

	intentKey, ok := request.URL.Query()["intent"]

	if !ok || len(intentKey[0]) < 1 {
		messageError := "Url Param 'intent' is missing"
		log.Println(messageError)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(messageError))
	} else {
		learning := mongoConnector.RetrieveLearning(intentKey[0])
		if learning == nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("No Learning Found"))
		} else {
			conversion, error := json.Marshal(learning)

			if error != nil {
				panic(error)
			}
			response.Write(conversion)
		}
	}
}

func searchInLearningDatabase(response http.ResponseWriter, request *http.Request) {
	userEntry := dtos.UserEntry{}
	agentResponse := dtos.AgentResponse{}
	
	error := json.NewDecoder(request.Body).Decode(&userEntry)

	if error != nil {
		panic(error)
	}

	matching := mongoConnector.Search(userEntry.Text)


	agentResponse.AgentResponse = "Lo siento, no te he entendio"
	agentResponse.UserSaid = userEntry.Text
	agentResponse.UserIntent = "Fallback"

	if matching != nil {
		agentResponse.UserIntent = matching.Intent
		agentResponse.AgentResponse = matching.AgentResponse[rand.Intn(len(matching.AgentResponse))]
	} 

	messageResponse, error := json.Marshal(agentResponse)

	if error != nil {
		panic(error)
	}
	response.Write(messageResponse)
	
}

func returnMessage(response http.ResponseWriter, request *http.Request, message string) {

	m := dtos.Message{message}
	bodyMessage, error := json.Marshal(m)

	if error != nil {
		panic(error)
	}

	response.Write(bodyMessage)
}
