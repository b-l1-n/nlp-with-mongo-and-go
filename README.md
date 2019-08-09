# nlp-with-mongo-and-go

We are trying to create a simple NLP based on Mongo and a Go Service.

For that purpose we use the following technologies:

* Docker: As an Awesome virtualization for services. A way to carry on your development to any computer.
* Go: As lightweigth programming language for all the agent logic, is a way to expose the knowledge for a consumer.
* Mongo: As a fake NLP using Text indexes for classify intents for users


If you want to run the agent, run the following commad:

```
docker-compose up -d
```

For retrieving the utterances of an specific intent, use the following command:

```
curl -X GET 'http://localhost:8080/agent/v1/learn?intent=anyIntentName'
```

If you want that the agent learn something, use:

```
curl -X POST \
  http://localhost:8080/agent/v1/learn \
  -H 'Content-Type: application/json' \
  -d '{
  "Intent": "hello",
  "Utterances": [
    "hola",
    "hola que tal",
    "ey",
    "oye"
  ],
  "AgentResponse": [
    "Hola!",
    "Muy buenas!"
  ],
  "AgentType" : "Trivial Agent"
}'
```


And finally, to talk with our agent use:

```
curl -X POST \
  http://localhost:8080/agent/v1/detectIntent \
  -d '{
        "Text": "holi" 
    }'
```

## Advanced Knowledge

In order to create entities that our agent can extract from the user utterance, we can add this information as follow:

```
curl -X POST \
  http://localhost:8080/agent/v1/entities \
  -d '{
  "Name" : "films",
  "Values" : [
  	"Django Unchained",
  	"Inception",
  	"Die Hard"
  ]
}'
```
