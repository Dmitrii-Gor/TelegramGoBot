package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func TextHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID

	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Привет! Отправь мне файл, docx/doc, а я верну его в формате pdf",
	})
}
