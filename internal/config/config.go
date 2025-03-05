package config

import (
	"fmt"
	"os"
)

var TelegramBotToken string

func Init() {
	TelegramBotToken = os.Getenv("TokenBot")
	if TelegramBotToken == "" {
		fmt.Println("Ошибка: переменная TokenBot не установлена")
		os.Exit(1)
	}
}
