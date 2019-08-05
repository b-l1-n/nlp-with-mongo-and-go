package agentKnowledge

import (
	"agent/dtos"
	"encoding/json"
	"net/http"
	"math/rand"
	"agent/database"
)

func DetectIntent(response http.ResponseWriter, request *http.Request) {
	userEntry := dtos.UserEntry{}
	agentResponse := dtos.AgentResponse{}
	
	error := json.NewDecoder(request.Body).Decode(&userEntry)

	if error != nil {
		panic(error)
	}

	matching := mongoConnector.Search(userEntry.Text)
	matching_entities := mongoConnector.ExtractEntities(userEntry.Text)


	agentResponse.AgentResponse = "Lo siento, no te he entendio"
	agentResponse.UserSaid = userEntry.Text
	agentResponse.UserIntent = "Fallback"

	if matching != nil {
		agentResponse.UserIntent = matching.Intent
		agentResponse.AgentResponse = matching.AgentResponse[rand.Intn(len(matching.AgentResponse))]
		agentResponse.AgentType = matching.AgentType
		if matching_entities != nil {
			agentResponse.Entities = matching_entities
		}
	} 

	messageResponse, error := json.Marshal(agentResponse)

	if error != nil {
		panic(error)
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(messageResponse)
	
}
