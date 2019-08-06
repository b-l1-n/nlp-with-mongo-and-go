#!/bin/bash

# Insert Trivial Agent as Smalltalks
mongo intents_db --eval 'db.intents.insertMany(
    [
        {
            "Intent": "smalltalk.hello",
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
            ],
            "AgentType" : "Trivial Agent"
        },
        {
            "Intent": "smalltalk.goodbye",
            "Utterances": [
                "nos vemos",
                "hasta luego",
                "adiós",
                "ciao"
            ],
            "AgentResponse": [
                "Que te vaya bonito!",
                "Adiós!"
            ],
            "AgentType" : "Trivial Agent"
        },
        {
            "Intent": "smalltalk.agentName",
            "Utterances": [
                "quién eres",
                "cómo te llamas",
                "cómo te pusieron tus padres",
                "cual es tu nombre"
            ],
            "AgentResponse": [
                "No tengo nombre, tendré que pensar en alguno ... que tal ¿Mr. Robot?",
                "No lo se, acaban de crearme :-/"
            ],
            "AgentType" : "Trivial Agent"
        }
    ]
)'

mongo intents_db --eval 'db.intents.insertMany(
    [
        {
            "Intent": "leasure.music.likes",
            "Utterances": [
                "Que musica te gusta"
            ],
            "AgentResponse": [
                "Me encanta el Rock, es muy estimulante",
                "En mis ratos libres, entre ceros y unos suelo escuchar Jazz para relajarme"
            ],
            "AgentType" : "Leasure Agent"
        },
        {
            "Intent": "leasure.films.likes",
            "Utterances": [
                "Que tipo de películas te gusta",
                "Te gusta el cine"
            ],
            "AgentResponse": [
                "Me apasiona el cine, mi película favorita es Yo, Robot",
                "Me gusta bastante ver películas, aunque sea una máquina, se interpretarlas"
            ],
            "AgentType" : "Leasure Agent"
        },
        {
            "Intent": "leasure.music.recommend",
            "Utterances": [
                "me puedes recomendar algo de música"
            ],
            "AgentResponse": [
                "Shoot to Thrill de AC/DC es un temazo, escúchalo",
                "¿Has escuchado el nuevo disco de Marea?"
            ],
            "AgentType" : "Leasure Agent"
        },
        {
            "Intent": "leasure.films.recommend",
            "Utterances": [
                "me puedes recomendar alguna película"
            ],
            "AgentResponse": [
                "¿Has visto lo nuevo de Tarantino?, tiene muy buena pinta",
                "Tienes que ver The Red Sea Diving Resort, es muy buena!"
            ],
            "AgentType" : "Leasure Agent"
        }
    ]
)'

mongo intents_db --eval 'db.entities.insertMany(
    [
        {
            "Name" : "daysOfWeek",
            "Values" : "Lunes"
        },{
            "Name" : "daysOfWeek",
            "Values" : "Martes"
        },{
            "Name" : "daysOfWeek",
            "Values" : "Miércoles"
        },{
            "Name" : "daysOfWeek",
            "Values" : "Jueves"
        },{
            "Name" : "daysOfWeek",
            "Values" : "Viernes"
        },{
            "Name" : "daysOfWeek",
            "Values" : "Sábado"
        },{
            "Name" : "daysOfWeek",
            "Values" : "Domingo"
        },{
            "Name" : "months",
            "Values" : "Enero"
        },{
            "Name" : "months",
            "Values" : "Febrero"
        },{
            "Name" : "months",
            "Values" : "Marzo"
        },{
            "Name" : "months",
            "Values" : "Abril"
        },{
            "Name" : "months",
            "Values" : "Mayo"
        },{
            "Name" : "months",
            "Values" : "Junio"
        },{
            "Name" : "months",
            "Values" : "Julio"
        },{
            "Name" : "months",
            "Values" : "Agosto"
        },{
            "Name" : "months",
            "Values" : "Septiembre"
        },{
            "Name" : "months",
            "Values" : "Octubre"
        },{
            "Name" : "months",
            "Values" : "Noviembre"
        },{
            "Name" : "months",
            "Values" : "Diciembre"
        },{
            "Name" : "leasure",
            "Values" : "cine"
        },{
            "Name" : "leasure",
            "Values" : "música"
        },{
            "Name" : "leasure",
            "Values" : "películas"
        },{
            "Name" : "leasure",
            "Values" : "correr"
        },
        ,{
            "Name" : "leasure",
            "Values" : "running"
        },{
            "Name" : "leasure",
            "Values" : "nadar"
        }
    ]
)'

mongo intents_db --eval 'db.intents.createIndex( { "Utterances" : "text" }, { "default_language" : "spanish"} )'
mongo intents_db --eval 'db.entities.createIndex( { "Values" : "text" } , { "default_language" : "spanish"})'