package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "app: ", log.Ldate|log.Ltime|log.Lshortfile)
	serv := server.NewServer(logger)

	if err := serv.ServerHttp.ListenAndServe(); err != nil {
		serv.Logger.Fatal(err)
	}
}
