FROM golang:1.8 as builder
ENV SRC=/go/app
RUN mkdir -p ${SRC}
ADD . ${SRC}
WORKDIR ${SRC}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/agent

FROM alpine:3.6
EXPOSE 8080
RUN mkdir /app
ENV WORKD_DIR=/app
WORKDIR ${WORKD_DIR}
COPY --from=builder /go/app/bin/agent ./
ENTRYPOINT ${WORKD_DIR}/agent