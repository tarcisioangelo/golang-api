package services

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken - JWT
func CreateToken(userID uint) (string, error) {
	var err error

	accessSecret := os.Getenv("ACCESS_SECRET")

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = userID
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(accessSecret))

	if err != nil {
		return "", err
	}

	return token, nil
}

// TokenValid - Validação do Token JWT
// Valida e retorna o ID do Usuário
func TokenValid(r string) (int, error) {
	tokenString := extractToken(r)

	accessSecret := os.Getenv("ACCESS_SECRET")

	type Claims struct {
		User int `json:"user"`
		jwt.StandardClaims
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(accessSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, err
	}

	return claims.User, nil
}

func extractToken(r string) string {
	bearToken := r

	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
