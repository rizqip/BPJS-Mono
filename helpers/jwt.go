package helpers

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func DecodeTokenJwt(tokenStr string) (jwt.MapClaims, bool) {

	hmacSecretString := os.Getenv("JWT_SECRET")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println("Invalid Jwt")
		return nil, false
	}
}
