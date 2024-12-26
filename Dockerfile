FROM golang:1.23.4

WORKDIR /api

COPY go.mod go.sum ./
COPY . .

RUN go mod tidy && go mod download && go mod verify

RUN go build -v -o gauth ./cmd
