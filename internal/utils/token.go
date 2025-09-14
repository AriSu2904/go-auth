package utils

import (
	"crypto/rsa"
	"github.com/AriSu2904/go-auth/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ExpiryTime struct {
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

func GenerateTokenJwt(user *models.User, privateKeys *rsa.PrivateKey, expiryTime *ExpiryTime) (*models.TokenInfo, error) {
	accessExpTime := time.Now().Add(expiryTime.AccessExpiry)
	accessClaims := &models.Claims{
		UserID:  user.ID,
		Role:    string(user.Role),
		Email:   user.Email,
		Persona: user.Persona,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpTime),
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.SigningMethodRS256, accessClaims)
	accessToken, err := accessJwt.SignedString(privateKeys)

	if err != nil {
		return &models.TokenInfo{}, err
	}

	refreshExpTime := time.Now().Add(expiryTime.RefreshExpiry)
	refreshClaims := &models.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpTime),
		},
	}

	refreshJwt := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshClaims)
	refreshToken, err := refreshJwt.SignedString(privateKeys)

	if err != nil {
		return &models.TokenInfo{}, err
	}

	return &models.TokenInfo{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
