package main

import (
	"log"
	"net/http"
	"./handlers"
)

func main() {
	log.Println("Starting server at 8080")
	http.HandleFunc("/agent/v1/detectIntent", agentHandler.DetectIntent)
	http.HandleFunc("/agent/v1/learn", agentHandler.LearnUtterance)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started....")
}
