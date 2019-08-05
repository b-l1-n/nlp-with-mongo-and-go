package dtos

type AgentResponse struct {
	AgentResponse string
	UserSaid string
	UserIntent string
	AgentType string
	Entities []*EntityDB
}

type UserEntry struct {
	Text string
}

type Learning struct {
	Intent string `bson:"Intent"`
	Utterances []string `bson:"Utterances"`
	AgentResponse []string `bson:"AgentResponse"`
	AgentType string `bson:"AgentType"`
}

type Message struct {
	Message string
}

type Entity struct {
	Name string
	Values []string
}

type EntityDB struct {
	Name string `bson:"Name"`
	Values string `bson:"Values"`
}
