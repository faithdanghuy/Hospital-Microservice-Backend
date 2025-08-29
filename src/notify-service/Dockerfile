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
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*
    
COPY --from=builder /app/notify-service/main .

EXPOSE 3004
CMD ["./main"]
