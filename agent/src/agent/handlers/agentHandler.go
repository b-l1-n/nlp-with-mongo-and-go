package agentHandler

import (
	"agent/dtos"
	"encoding/json"
	"net/http"
	"agent/services"
)

func DetectIntent(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		agentKnowledge.DetectIntent(response, request)
	default:
		returnMessage(response, request, "Method not allowed")
	}
}

func LearnUtterance(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		agentKnowledge.InsertIntoLearningDatabase(response, request)
	case "GET":
		agentKnowledge.RetrieveLearningFromDatabase(response, request)
	default:
		returnMessage(response, request, "Method not allowed")
	}
}

func EntitiesManager(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		agentKnowledge.CreateNewEntity(response, request)
	case "GET":
		agentKnowledge.RetrieveAllAvailableEntities(response, request)
	default:
		returnMessage(response, request, "Method not allowed")
	}
}

func returnMessage(response http.ResponseWriter, request *http.Request, message string) {

	m := dtos.Message{message}
	bodyMessage, error := json.Marshal(m)

	if error != nil {
		panic(error)
	}

	response.Write(bodyMessage)
}
