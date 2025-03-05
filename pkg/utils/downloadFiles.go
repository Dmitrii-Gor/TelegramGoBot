package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(fileName, url string) error {
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
