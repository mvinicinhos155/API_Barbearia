package middleware

import (
	"api_barbearia/internal/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "token obrigatorio", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "token mal formatado", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritmo inválido")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		ctx := context.WithValue(r.Context(), "userID", claims.ID)

			if err != nil {
				fmt.Println("JWT ERROR:", err)
				http.Error(w, "token invalido", http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				http.Error(w, "token invalido", http.StatusUnauthorized)
				return
			}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}