package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"work.com/pkg/application"
)

//* Загрузка файла окружения, в нём хранятся ключи шифрования токенов
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file.")
	}
}

func main() {
	app := application.NewApplication()

	s := http.Server{
		Handler:  app.Handler,
		ErrorLog: app.ErrLog,
		Addr:     app.Config.Addr,
	}

	app.Info("Server is starting...")
	app.Error(s.ListenAndServe())
}
