FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY apps/bot-sma ./apps/bot-sma
COPY pkg ./pkg
COPY api ./api

RUN go install github.com/air-verse/air@latest

RUN go build -o bot-sma ./apps/bot-sma
CMD ["./bot-sma"]
