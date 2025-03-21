package main

import (
	"TelegramGoBot/internal/config"
	"TelegramGoBot/internal/handlers"
	"context"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

// Send any text message to the bot after the bot has been started

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	config.Init()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.HandlerRoute),
	}

	b, err := bot.New(config.TelegramBotToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}
