package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HandlerRoute(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch {
	case update.Message != nil && update.Message.Document != nil:
		DocumentHandler(ctx, b, update)
	case update.Message != nil && update.Message.Text != "":
		TextHandler(ctx, b, update)
	default:
		unknownHandler(ctx, b, update)
	}
}

func unknownHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Я не понимаю этот формат Отправьте текст или документ.",
	})
}
