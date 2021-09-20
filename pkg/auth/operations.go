package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"work.com/pkg/storage"
)

func SendTokenPair(w http.ResponseWriter, userid string, db *storage.DB) (int, error) {
	//* Генерация новой пары токенов
	at, rt, err := GenerateNewTokens(userid)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	//* Упаковка токенов в Cookies
	http.SetCookie(w, &http.Cookie{
		Name:  "Access-Token",
		Value: base64.RawStdEncoding.EncodeToString([]byte(at)),
		// Secure:   true, 
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "Refresh-Token",
		Value: base64.RawStdEncoding.EncodeToString([]byte(rt)),
		// Secure:   true, 
		HttpOnly: true,
	})

	//* Хеширование Refresh-токена и сохранение в базе данных
	hashedRt, err := bcrypt.GenerateFromPassword(
		[]byte(rt),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	db.Update(userid, string(hashedRt))
	return http.StatusOK, nil
}

func RefreshTokenPair(w http.ResponseWriter, r *http.Request, userid string, db *storage.DB) (int, error) {
	//* Извлечение токенов из Cookie в формате Base64
	baseAt, err := r.Cookie("Access-Token")
	if err != nil {
		return http.StatusBadRequest, ErrEmptyToken
	}
	baseRt, err := r.Cookie("Refresh-Token")
	if err != nil {
		return http.StatusBadRequest, ErrEmptyToken
	}

	//* Расшифровка токенов
	at, err := base64.RawStdEncoding.DecodeString(baseAt.Value)
	if err != nil {
		return http.StatusBadRequest, ErrUnableToDecodeToken
	}

	rt, err := base64.RawStdEncoding.DecodeString(baseRt.Value)
	if err != nil {
		return http.StatusBadRequest, ErrUnableToDecodeToken
	}

	//* Получение хеша Refresh-токена из базы данных и сопоставление
	//* его с полученным токеном
	encRt, err := db.Get(userid)
	if err != nil {
		return http.StatusBadRequest, ErrNoTokenInDatabase
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(fmt.Sprintf("%v", encRt.RefreshToken)),
		[]byte(rt),
	); err != nil {
		return http.StatusUnauthorized, ErrInvalidToken
	}

	//* Проверка пары на валидность
	atUserid, atTimestamp, err := ParseToken(string(at), "ACCESS_KEY")
	if err != nil {
		return http.StatusUnauthorized, err
	}

	rtUserid, rtTimestamp, err := ParseToken(string(rt), "REFRESH_KEY")
	if err != nil {
		return http.StatusUnauthorized, err
	}

	if rtUserid != atUserid || rtTimestamp != atTimestamp {
		return http.StatusUnauthorized, ErrInvalidTokenPair
	}

	//* В случае валидности пары отправляем новую пару пользователю
	statusCode, err := SendTokenPair(w, userid, db)
	if err != nil {
		return statusCode, err
	}

	return http.StatusOK, nil
}
