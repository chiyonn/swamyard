FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o aggregator ./apps/aggregator

CMD ["./aggregator"]
