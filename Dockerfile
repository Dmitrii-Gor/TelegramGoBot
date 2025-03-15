FROM golang:1.24-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o telegram-bot ./cmd/bot/main.go

CMD ["./telegram-bot"]