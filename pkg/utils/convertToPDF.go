package utils

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WordToPdfConvert(ctx context.Context, fileName string) (string, error) {

	tokenSource, err := GetGoogleTokenData(ctx)
	if err != nil {
		log.Fatalf("Ошибка получения TokenSource: %v", err)
	}

	srv, err := drive.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Fatalf("Ошибка создания сервиса Drive: %v", err)
	}

	// Открываем DOCX файл
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Параметры загрузки файла
	fileMetadata := &drive.File{
		Name:     fileName,         // Название файла в Drive
		Parents:  []string{"root"}, // Загружаем в корневую папку
		MimeType: "application/vnd.google-apps.document",
	}

	// Загружаем файл в Google Drive
	uploadedFile, err := srv.Files.Create(fileMetadata).
		Media(file).
		Do()
	if err != nil {
		log.Fatalf("Ошибка загрузки: %v", err)
	}

	fmt.Printf("Файл загружен, ID: %s\n", uploadedFile.Id)

	// Скачиваем документ в формате PDF
	exportFileID := uploadedFile.Id
	exportMimeType := "application/pdf"
	exportedFile, err := srv.Files.Export(exportFileID, exportMimeType).Download()
	if err != nil {
		log.Fatalf("Ошибка экспорта: %v", err)
	}

	// Сохраняем в локальный файл
	extension := filepath.Ext(fileName)
	fileNameWithoutExtension := strings.TrimSuffix(fileName, extension)
	newPdfFile := fmt.Sprintf("%s.pdf", strings.TrimSuffix(fileNameWithoutExtension, ".docx"))
	outFile, err := os.Create(newPdfFile)
	if err != nil {
		log.Fatalf("Ошибка создания файла: %v", err)
	}
	defer outFile.Close()

	_, err = outFile.ReadFrom(exportedFile.Body)
	if err != nil {
		log.Fatalf("Ошибка записи в файл: %v", err)
	}

	err = srv.Files.Delete(uploadedFile.Id).Do()
	if err != nil {
		log.Fatalf("Ошибка удаления файла с drive: %v", err)
	}

	return newPdfFile, nil
}
