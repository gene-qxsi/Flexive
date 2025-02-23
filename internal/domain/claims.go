package domain

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}
