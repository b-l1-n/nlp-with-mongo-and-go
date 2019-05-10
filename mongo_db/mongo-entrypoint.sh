#!/bin/bash


mongo intents_db --eval 'db.intents.insertMany(
    [
        {
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
            ]
        },
        {
            "Intent": "goodbye",
            "Utterances": [
                "nos vemos",
                "hasta luego",
                "adiós",
                "ciao"
            ],
            "AgentResponse": [
                "Que te vaya bonito!",
                "Adiós!"
            ]
        }
    ]
)'
mongo intents_db --eval 'db.intents.createIndex( { "Utterances" : "text" } )'