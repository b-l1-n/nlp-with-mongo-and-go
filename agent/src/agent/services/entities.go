package agentKnowledge

import (
	"agent/dtos"
	"encoding/json"
	"net/http"
	"agent/database"
)

func CreateNewEntity(response http.ResponseWriter, request *http.Request) {
	entity := dtos.Entity{}

	error := json.NewDecoder(request.Body).Decode(&entity)

	if error != nil {
		panic(error)
	}

	ok := mongoConnector.InsertEntity(entity)

	if ok {
		conversion, error := json.Marshal(entity)

		if error != nil {
			panic(error)
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(conversion)
	}

	response.Write(nil)
}

func RetrieveAllAvailableEntities(response http.ResponseWriter, request *http.Request) {

	entities := mongoConnector.RetrieveEntities()
	if entities == nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("No Entities Found"))
	} else {
		conversion, error := json.Marshal(entities)

		if error != nil {
			panic(error)
		}
		response.Header().Set("Content-Type", "application/json")
		response.Write(conversion)
	}

}
