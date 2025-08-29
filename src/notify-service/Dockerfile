# 1. Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY ../hospital-core ./hospital-core

RUN go mod download

COPY . .

WORKDIR /app/notify-service
RUN go build -o main ./cmd/app

# 2. Runtime stage
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/notify-service/main .
COPY notify-service/deploy/ca.pem ./ca.pem

EXPOSE 3002
CMD ["./main"]
