FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o bot-sma ./apps/bot-sma

CMD ["./bot-sma"]
