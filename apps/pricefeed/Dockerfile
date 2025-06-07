FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY apps/pricefeed ./apps/pricefeed
COPY pkg ./pkg
COPY api ./api

RUN go install github.com/air-verse/air@latest

RUN go build -o pricefeed ./apps/pricefeed

EXPOSE 8080
EXPOSE 50052

CMD ["./pricefeed"]
