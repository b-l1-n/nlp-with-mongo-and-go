package dtos

type AgentResponse struct {
	AgentResponse string
	UserSaid string
	UserIntent string
}

type UserEntry struct {
	Text string
}

type Learning struct {
	Intent string `bson:"Intent"`
	Utterances []string `bson:"Utterances"`
	AgentResponse []string `bson:"AgentResponse"`
}

type Message struct {
	Message string
}
