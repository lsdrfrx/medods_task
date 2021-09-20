package application

import "os"

//* Структура, необходимая для конфигурации приложения
type Config struct {
	//* Адрес, на котором запустится сервер
	Addr string
}

func NewConfig() *Config {
	return &Config{
		Addr: os.Getenv("ADDR"),
	}
}
