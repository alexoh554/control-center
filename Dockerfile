FROM golang:1.21-alpine

RUN apk add --no-cache git bash

WORKDIR /app

# Copy go modules
COPY go.mod /
RUN go mod download

COPY modules/ ./modules
