package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Это наш собственный сервер — просто обёртка над http.Server
type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

// Функция New создаёт новый сервер
func New(logger *log.Logger) *Server {
	// Создаём стандартный маршрутизатор из net/http
	mux := http.NewServeMux()

	// Регистрируем обработчики
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	// Настраиваем параметры сервера
	server := &http.Server{
		Addr:         ":8080",          // порт
		Handler:      mux,              // маршруты
		ErrorLog:     logger,           // логгер ошибок
		ReadTimeout:  5 * time.Second,  // время на чтение запроса
		WriteTimeout: 10 * time.Second, // время на отправку ответа
		IdleTimeout:  15 * time.Second, // время между запросами
	}

	// Возвращаем нашу структуру Server
	return &Server{
		httpServer: server,
		logger:     logger,
	}
}

// Start запускает сервер
func (s *Server) Start() error {
	s.logger.Println("Мана есть! (Сервер запущен на порту :8080)")
	return s.httpServer.ListenAndServe()
}
