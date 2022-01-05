# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR cache/

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go get -u google.golang.org/grpc

COPY *.go ./
COPY *.json ./

RUN go build -o /cache

EXPOSE 8080

CMD [ "/cache" ]