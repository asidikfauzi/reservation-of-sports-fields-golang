package middlewares

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"
)

func CreateToken(userId string) (string, error) {
	var appConfig map[string]string

	appConfig, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error reading .env file")
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(appConfig["SECRET_JWT"]))
}

func ExtractTokenUser(c *gin.Context) int {
	user, exists := c.Get("user")
	if !exists {
		return 0
	}
	token, ok := user.(*jwt.Token)
	if !ok {
		return 0
	}
	if !token.Valid {
		return 0
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}
	userId, ok := claims["userId"].(int)
	if !ok {
		return 0
	}
	return userId
}
