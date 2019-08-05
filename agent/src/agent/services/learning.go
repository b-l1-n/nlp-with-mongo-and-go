package agentKnowledge

import (
	"agent/dtos"
	"encoding/json"
	"log"
	"net/http"
	"agent/database"
)

func InsertIntoLearningDatabase(response http.ResponseWriter, request *http.Request) {
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
		response.Header().Set("Content-Type", "application/json")
		response.Write(conversion)
	}

	response.Write(nil)
}

func RetrieveLearningFromDatabase(response http.ResponseWriter, request *http.Request) {

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
			response.Header().Set("Content-Type", "application/json")
			response.Write(conversion)
		}
	}
}
