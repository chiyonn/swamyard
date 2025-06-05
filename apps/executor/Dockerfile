FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY apps/executor ./apps/executor
COPY pkg ./pkg
COPY api ./api

RUN go install github.com/air-verse/air@latest

RUN go build -o executor ./apps/executor
CMD ["./executor"]
