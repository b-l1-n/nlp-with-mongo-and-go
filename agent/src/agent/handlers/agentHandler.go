package agentHandler

import (
	"encoding/json"
	"net/http"

	"agent/database"
	"agent/dtos"
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

	ok := mongoConnector.Insert(learning)

	if ok {
		conversion, error := json.Marshal(learning)

		if error != nil {
			panic(error)
		}

		response.Write(conversion)
	}

	response.Write(nil)
}

func searchInLearningDatabase(response http.ResponseWriter, request *http.Request) {
	userEntry := dtos.UserEntry{}

	error := json.NewDecoder(request.Body).Decode(&userEntry)

	if error != nil {
		panic(error)
	}

	matching := mongoConnector.Search(userEntry.Text)

	messageResponse, error := json.Marshal(matching)

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
