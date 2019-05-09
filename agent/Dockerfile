FROM golang:1.8 as builder
RUN go get -u github.com/golang/dep/cmd/dep
ENV SRC=/go/src
RUN mkdir -p ${SRC}
ADD src ${SRC}
WORKDIR ${SRC}/agent
RUN rm Gopkg.toml
RUN rm Gopkg.lock
RUN dep init
RUN dep ensure -add "go.mongodb.org/mongo-driver/mongo"
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/agent
RUN ls -la ${SRC}/agent/bin

FROM alpine:3.6
EXPOSE 8080
RUN mkdir /app
ENV WORKD_DIR=/app
WORKDIR ${WORKD_DIR}
COPY --from=builder /go/src/agent/bin/agent ./
ENTRYPOINT ${WORKD_DIR}/agent