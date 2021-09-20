package auth

import "github.com/golang-jwt/jwt"

//* Структура, необходимая для записи пользовательской информации
//* в токены и получения её из них
type Claims struct {
	jwt.StandardClaims
	Userid    string `json:"userid"`
	Timestamp int64  `json:"timestamp"`
}
