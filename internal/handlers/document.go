package handlers

import (
	"TelegramGoBot/internal/config"
	"TelegramGoBot/pkg/utils"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
)

func DocumentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: update.Message.Document.FileID,
	})
	if err != nil {
		fmt.Println("Ошибка получения файла:", err)
		return
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", config.TelegramBotToken, file.FilePath)

	localFilePath := update.Message.Document.FileName
	err = utils.DownloadFile(localFilePath, fileURL)
	if err != nil {
		fmt.Println("Ошибка скачивания файла:", err)
	}

	pfdFileToReturn, err := utils.WordToPdfConvert(ctx, localFilePath)
	if err != nil {
		fmt.Println("Ошибка в конвертации файла:", err)
		return
	}

	fileData, err := os.Open(pfdFileToReturn)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer func() {
		fileData.Close()
		err = os.Remove(localFilePath)
		err = os.Remove(pfdFileToReturn)
		if err != nil {
			fmt.Println("Ошибка удаления файла:", err)
		}
	}()

	// Отправляем файл
	_, err = b.SendDocument(ctx, &bot.SendDocumentParams{
		ChatID: update.Message.Chat.ID,
		Document: &models.InputFileUpload{
			Filename: pfdFileToReturn,
			Data:     fileData,
		},
	})
	if err != nil {
		fmt.Println("Ошибка отправки файла:", err)
		return
	}

}
