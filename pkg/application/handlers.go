package application

import (
	"fmt"
	"net/http"

	"work.com/pkg/auth"
)

//* Маршрут получения пары токенов
func Recieve(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//* Получение UserID и проверка его наличия
		userid := r.URL.Query().Get("userid")
		if userid == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			app.Error("Username is empty")
			return
		}

		//* Проверка на наличие пользователя в базе
		//* и добавление новой записи в случае
		//* отсутствия
		val, _ := app.Database.Get(userid)
		if val.UserId != "" {
			http.Error(w, "User already exists", http.StatusBadRequest)
			app.Error("User already exists")
			return
		}
		if err := app.Database.Create(userid); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			app.Error("Unable to create new document:", err.Error())
			return
		}

		//* Генерация и отправка пользователю пары токенов
		statusCode, err := auth.SendTokenPair(w, userid, app.Database)
		if err != nil {
			http.Error(w, http.StatusText(statusCode), statusCode)
			app.Error(err.Error())
			return
		}

		fmt.Fprintf(w, "Tokens recieved successful")
	}
}

//* Маршрут обновления пары токенов
func Refresh(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//* Получение UserID и проверка его наличия
		userid := r.URL.Query().Get("userid")
		if userid == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			app.Error("Username is empty")
			return
		}

		//* Проверка пары токенов на валидность, проведение Refresh-операции
		statusCode, err := auth.RefreshTokenPair(w, r, userid, app.Database)
		if err != nil {
			http.Error(w, http.StatusText(statusCode), statusCode)
			app.Error(err.Error())
			return
		}

		fmt.Fprintf(w, "Tokens refreshed successful")
	}
}
