FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR cmd/hackathon-tg-bot

RUN go build
