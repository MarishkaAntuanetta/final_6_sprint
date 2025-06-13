package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаём стандартный логгер для вывода сообщений в консоль
	logger := log.Default()

	// Вызываем функцию New из пакета server, передаем логгер
	srv := server.New(logger)

	// Пишем в лог, что собираемся запустить сервер
	logger.Println("Генерирую ману (Запускаю сервер на порту :8080)")

	// Запускаем сервер
	err := srv.Start()
	if err != nil {
		// Если есть ошибка — пишем в лог и завершаем программу
		logger.Fatal("Не хватило маны. Сбой в работе сервера", err)
	}
}
