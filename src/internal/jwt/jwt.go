package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"os"
	"time"
)

func GenerateToken(username string, isAdmin bool) string {
	role := "employee"
	if isAdmin {
		role = "admin"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Errorf("Error generating token: %s", err)
	}

	return tokenString
}
