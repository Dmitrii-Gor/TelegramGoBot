package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"io"
	"net/http"
	"os"
	"os/signal"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("YOUR_BOT_TOKEN", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println(update.Message.Document.FileID)
	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: update.Message.Document.FileID,
	})
	if err != nil {
		fmt.Println("Ошибка получения файла:", err)
		return
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", "YOUR_BOT_TOKEN", file.FilePath)
	fmt.Println("Ссылка на скачивание:", fileURL)

	downloadFile(update.Message.Document.FileName, fileURL)
}

// Функция для скачивания файла
func downloadFile(fileName, url string) error {
	// Делаем HTTP-запрос к Telegram-серверу
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаём файл на диске
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Копируем данные из HTTP-ответа в файл
	_, err = io.Copy(out, resp.Body)
	return err
}
