package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler — просто открывает index.html и отправляет браузеру
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Открываем файл index.html
	http.ServeFile(w, r, "index.html")
}

// UploadHandler — принимает файл, конвертирует и сохраняет результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму (максимум 10 МБ)
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Ошибка: не удалось распосзнать форму", http.StatusInternalServerError)
		return
	}
	// Получаем файл по имени myFile
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка: не удалось получить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка: не удалось прочитать файл", http.StatusInternalServerError)
		return
	}

	// Проверяем, что файл не пустой
	if len(data) == 0 {
		http.Error(w, "Ошибка: файл пустой", http.StatusInternalServerError)
		return
	}

	// Передаём данные в сервис для конвертации
	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Ошибка: конвертация не удалась", http.StatusInternalServerError)
		return
	}

	// Создаём имя файла с текущей датой
	filename := fmt.Sprintf("%s.txt", time.Now().UTC().Format("20060102-150405"))
	outputPath := filepath.Join("results", filename)

	// Создаём папку results, если её нет
	err = os.MkdirAll("results", os.ModePerm)
	if err != nil {
		http.Error(w, "Не удалось создать папку", http.StatusInternalServerError)
		return
	}

	// Создаём файл для записи результата
	outputFile, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Не удалось создать файл", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ пользователю
	response := fmt.Sprintf("Вы загрузили: %s\n\nРезультат:\n%s", string(data), result)
	_, err = w.Write([]byte(response))
	if err != nil {
		log.Printf("Ошибка при отправке ответа клиенту: %v", err)
	}
}
