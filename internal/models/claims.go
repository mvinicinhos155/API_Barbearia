package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.StandardClaims
}