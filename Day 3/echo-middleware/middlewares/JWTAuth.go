package middlewares

import (
	"middleware/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = config.JWTExp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)
		return int(userId)
	}
	return 0
}
