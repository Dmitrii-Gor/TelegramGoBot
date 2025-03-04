package handlers

import (
	"TelegramGoBot/internal/config"
	"TelegramGoBot/pkg/utils"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: update.Message.Document.FileID,
	})
	if err != nil {
		fmt.Println("Ошибка получения файла:", err)
		return
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", config.TelegramBotToken, file.FilePath)
	fmt.Println("Ссылка на скачивание:", fileURL)

	err = utils.DownloadFile(update.Message.Document.FileName, fileURL)
	if err != nil {
		fmt.Println("Ошибка скачивания файла:", err)
	}

}
