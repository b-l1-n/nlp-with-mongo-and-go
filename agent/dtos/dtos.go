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
	Intent string
	Utterances []string
	AgentResponse []string
}

type Message struct {
	Message string
}
