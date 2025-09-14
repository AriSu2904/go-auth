package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID  string `json:"userId"`
	Email   string `json:"email"`
	Persona string `json:"persona"`
	Role    string `json:"role"`
	Issuer  string `json:"issuer" default:"go-auth"`
	jwt.RegisteredClaims
}

type TokenInfo struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
