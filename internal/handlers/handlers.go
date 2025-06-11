package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler — просто открывает index.html и отправляет браузеру
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что запрашивается именно главная страница
	if r.URL.Path != "/" {
		http.Error(w, "404 Страница не найдена", http.StatusNotFound)
		return
	}

	// Открываем файл index.html
	file, err := os.Open("../index.html")
	if err != nil {
		http.Error(w, "Главная страница недоступна", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html")

	// Копируем содержимое файла в ответ
	io.Copy(w, file)
}

// UploadHandler — принимает файл, конвертирует и сохраняет результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму (максимум 10 МБ)
	r.ParseMultipartForm(10 << 20)

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
		http.Error(w, "Ошибка: файл пустой", http.StatusBadRequest)
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
		http.Error(w, "Не могу создать папку", http.StatusInternalServerError)
		return
	}

	// Создаём файл для записи результата
	outputFile, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Не могу создать файл", http.StatusInternalServerError)
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
	fmt.Fprintf(w, "Готово!\n\nРезультат:\n%s", result)
}
