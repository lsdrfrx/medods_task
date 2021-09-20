package application

import (
	"log"
	"net/http"
	"os"

	"work.com/pkg/storage"
)

//* Структура, связывающая все компоненты приложения
type Application struct {
	Database *storage.DB
	Handler  *http.ServeMux
	ErrLog   *log.Logger
	InfoLog  *log.Logger
	Config   *Config
}

func NewApplication() *Application {
	app := &Application{
		Database: storage.NewDB(),
		ErrLog:   log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime),
		Config:   NewConfig(),
	}

	app.configureRouter()
	err := app.Database.Open()
	if err != nil {
		app.Error(err.Error())
	}

	return app
}

//* Настройка маршрутизатора приложения
func (app *Application) configureRouter() {
	router := http.NewServeMux()

	router.HandleFunc("/recieve", Recieve(app))
	router.HandleFunc("/refresh", Refresh(app))

	app.Handler = router
}

//* Хелпер для вывода информации в поток Stderr
func (app *Application) Error(content ...interface{}) {
	app.ErrLog.Println(content...)
}

//* Хелпер для вывода информации в поток Stdout
func (app *Application) Info(content ...interface{}) {
	app.InfoLog.Println(content...)
}