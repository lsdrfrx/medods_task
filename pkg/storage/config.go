package storage

//* Структура, необходимая для конфигурирования базы данных
type Config struct {
	URI            string
	DatabaseName   string
	CollectionName string
}

func NewConfig() *Config {
	return &Config{
		URI:            "mongodb://localhost:27017",
		DatabaseName:   "test",
		CollectionName: "test",
	}
}
