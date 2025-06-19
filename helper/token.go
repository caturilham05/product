package helper

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secret []byte

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		panic("Error loading .env file")
	}

	secret = []byte(os.Getenv("ENCRYPTED_SECRET_KEY"))
}

// func init() {
// 	// Load .env saat package helper diinisialisasi
// 	err := godotenv.Load()
// 	if err != nil {
// 		panic("Error loading .env file")
// 	}

// 	// Ambil SECRET_KEY dari .env
// 	secret = []byte(os.Getenv("ENCRYPTED_SECRET_KEY"))
// }

type JWTClaims struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	IsRefresh bool   `json:"is_refresh"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	if claims.IsRefresh {
		return nil, errors.New("Unauthorized")
	}

	return claims, nil
}
