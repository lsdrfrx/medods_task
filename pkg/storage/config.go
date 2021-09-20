package storage

import "os"

//* Структура, необходимая для конфигурирования базы данных
type Config struct {
	URI            string
	DatabaseName   string
	CollectionName string
}

func NewConfig() *Config {
	return &Config{
		URI:            os.Getenv("DATABASE_URI"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
	}
}
