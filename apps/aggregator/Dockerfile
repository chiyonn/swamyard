FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY apps/aggregator ./apps/aggregator
COPY pkg ./pkg
COPY api ./api

RUN go install github.com/air-verse/air@latest

RUN go build -o aggregator ./apps/aggregator
CMD ["./aggregator"]
