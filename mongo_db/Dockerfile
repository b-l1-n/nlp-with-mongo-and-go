FROM mongo:3.4
ENV MONGO_INITDB_DATABASE=intents_db
ADD mongo-entrypoint.sh /docker-entrypoint-initdb.d
RUN chmod +x /docker-entrypoint-initdb.d/mongo-entrypoint.sh