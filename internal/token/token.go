package token

import (
	"os"
	"time"
	"api_barbearia/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt (user *models.Users) (string, error) {

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	claims := models.Claims{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			return "", err
		}

	return  tokenString, nil
}