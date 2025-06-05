FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go build -o bot-sma ./apps/bot-sma

CMD ["./bot-sma"]
