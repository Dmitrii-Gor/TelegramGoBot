# Этап 1: Сборка приложения
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o telegram-bot ./cmd/bot/main.go

# Этап 2: Финальный образ
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/telegram-bot .
COPY token.json credentials.json .env ./
RUN apk --no-cache add ca-certificates

# Используем ENTRYPOINT вместо CMD
ENTRYPOINT ["./telegram-bot"]