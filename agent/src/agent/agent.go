package main

import (
	"log"
	"net/http"
	"agent/handlers"
)

func main() {
	log.Println("Starting server at 8080")
	http.HandleFunc("/agent/v1/detectIntent", agentHandler.DetectIntent)
	http.HandleFunc("/agent/v1/learn", agentHandler.LearnUtterance)
	http.HandleFunc("/agent/v1/entities", agentHandler.EntitiesManager)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started....")
}
