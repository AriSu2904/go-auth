package config

import (
	"crypto/rsa"
	"github.com/AriSu2904/go-auth/internal/utils"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	DBSource              string
	PrivateKey            *rsa.PrivateKey
	PublicKey             *rsa.PublicKey
	JwtIssuer             string
	JwtAccessTokenExpiry  time.Duration
	JwtRefreshTokenExpiry time.Duration
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file", err)
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE is not set")
	}

	jwtIssuer := os.Getenv("JWT_ISSUER")
	jwtAccessTokenExpiry := os.Getenv("JWT_ACCESS_TOKEN_EXPIRY")
	jwtRefreshTokenExpiry := os.Getenv("JWT_REFRESH_TOKEN_EXPIRY")

	if jwtIssuer == "" || jwtAccessTokenExpiry == "" || jwtRefreshTokenExpiry == "" {
		log.Fatal("JWT_ISSUER or JWT_ACCESS_TOKEN_EXPIRY or JWT_REFRESH_TOKEN_EXPIRY is not set")
	}

	privateKey := os.Getenv("JWT_PRIVATE_KEY")
	publicKey := os.Getenv("JWT_PUBLIC_KEY")

	if privateKey == "" || publicKey == "" {
		log.Fatal("JWT_PRIVATE_KEY or JWT_PUBLIC_KEY is not set")
	}

	loadedPrivateKey, err := utils.LoadPrivateKey(privateKey)
	if err != nil {
		log.Fatal("Failed to load private key:", err)
	}

	loadedPublicKey, err := utils.LoadPublicKey(publicKey)
	if err != nil {
		log.Fatal("Failed to load public key:", err)
	}

	parsedAccessExpiry, err := time.ParseDuration(jwtAccessTokenExpiry)
	if err != nil {
		log.Fatal("Failed to parse JWT_ACCESS_TOKEN_EXPIRY:", err)
	}

	parsedRefreshTokenExpiry, err := time.ParseDuration(jwtRefreshTokenExpiry)

	if err != nil {
		log.Fatal("Failed to parse JWT_REFRESH_TOKEN_EXPIRY:", err)
	}

	return &Config{
		DBSource:              dbSource,
		PrivateKey:            loadedPrivateKey,
		PublicKey:             loadedPublicKey,
		JwtIssuer:             jwtIssuer,
		JwtAccessTokenExpiry:  parsedAccessExpiry,
		JwtRefreshTokenExpiry: parsedRefreshTokenExpiry,
	}, nil
}
