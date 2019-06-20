#!/bin/bash


mongo intents_db --eval 'db.intents.insertMany(
    [
        {
            "Intent": "hello",
            "Utterances": [
                "hola",
                "hola que tal",
                "ey",
                "oye",
                "holi"
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
        },
        {
            "Intent": "agentName",
            "Utterances": [
                "quién eres",
                "cómo te llamas",
                "cómo te pusieron tus padres",
                "cual es tu nombre"
            ],
            "AgentResponse": [
                "No tengo nombre, tendré que pensar en alguno ... que tal ¿Mr. Robot?",
                "No lo se, acaban de crearme :-/"
            ]
        }
    ]
)'
mongo intents_db --eval 'db.intents.createIndex( { "Utterances" : "text" } )'