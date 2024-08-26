package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
