package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger     *log.Logger
	ServerHttp *http.Server
}

func NewServer(l *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.MainHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	serv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     l,
		ServerHttp: serv,
	}
}
