package config

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"

	"log"
	"os"
)

// Загружает токен из файла
func tokenFromFile(filename string) (*oauth2.Token, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Сохраняет токен в файл
func saveToken(filename string, token *oauth2.Token) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func GetGoogleTokenData(ctx context.Context) (oauth2.TokenSource, error) {
	// Загружаем креденшиалы (требуется, если надо обновлять токен)
	credentialsFile := "credentials.json"
	credentialsData, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Ошибка загрузки креденшиалов: %v", err)
	}

	// Создаём конфиг OAuth2
	config, err := google.ConfigFromJSON(credentialsData, drive.DriveScope)
	if err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	// Загружаем токен из файла
	token, err := tokenFromFile("token.json")
	if err != nil {
		log.Fatalf("Ошибка загрузки токена: %v", err)
	}

	// Обновляем токен при необходимости
	tokenSource := config.TokenSource(ctx, token)
	newToken, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Ошибка обновления токена: %v", err)
	}

	// Если токен обновился, сохраняем его
	if newToken.AccessToken != token.AccessToken {
		if err := saveToken("token.json", newToken); err != nil {
			return nil, fmt.Errorf("ошибка сохранения токена: %v", err)
		}
		fmt.Println("Токен обновлён и сохранён!")
	}

	return tokenSource, nil
}
