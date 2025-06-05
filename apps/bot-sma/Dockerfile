FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go build -o bot-sma ./cmd
CMD ["./bot-sma"]
HEALTHCHECK --interval=10s --timeout=2s CMD curl -f http://localhost:8080/health || exit 1
