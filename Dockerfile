FROM golang:1.21-alpine

RUN apk add --no-cache git bash

WORKDIR /app

# Go modules
COPY go.mod /
RUN go mod download

COPY modules/ ./modules
COPY migrations/ ./migrations

# Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
