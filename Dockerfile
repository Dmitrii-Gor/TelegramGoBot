# Этап 1: Сборка приложения
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы модулей и зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -v -o telegram-bot ./cmd/bot/main.go

# Этап 2: Финальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарник из этапа сборки
COPY --from=builder /app/telegram-bot .

# Копируем необходимые файлы (token.json, credentials.json, .env)
COPY token.json credentials.json .env ./

# Устанавливаем зависимости для работы с Google API (если нужно)
RUN apk --no-cache add ca-certificates

# Команда для запуска приложения
CMD ["./telegram-bot"]