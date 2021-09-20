package auth

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

//* Генерация пары токенов
//* Timestamp служит для дополнительного связывания Access и Refresh токенов
func GenerateNewTokens(userid string) (string, string, error) {
	timestamp := time.Now()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		Userid:    userid,
		Timestamp: timestamp.Unix(),
		StandardClaims: jwt.StandardClaims{
			//* Время жизни Access-токена - 5 минут
			ExpiresAt: timestamp.Add(5 * time.Minute).Unix(),
		},
	})

	rt := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		Userid: userid,
		Timestamp: timestamp.Unix(),
		StandardClaims: jwt.StandardClaims{
			//* Время жизни Refresh-токена - 24 часа
			ExpiresAt: timestamp.Add(24 * time.Hour).Unix(),
		},
	})	

	accessToken, err := at.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := rt.SignedString([]byte(os.Getenv("REFRESH_KEY")))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

//* Валидация токенов
//* Возвращает UserID и Timestamp из Claims валидного токена
//* В случае с невалидным Access токеном функция также возвращает
//* UserID и Timestamp для проведения Refresh-операции
func ParseToken(authToken, key string) (string, int64, error) {
	token, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSingingMethod
		}

		return []byte(os.Getenv(key)), nil
	})

	if err != nil && !strings.Contains(err.Error(), "expired") {
		return "", 0, err
	}

	if claims, ok := token.Claims.(*Claims); (key == "ACCESS_KEY" && ok) || token.Valid {
		return claims.Userid, claims.Timestamp, nil
	}

	return "", 0, ErrInvalidToken
}
