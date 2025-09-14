package service

import (
	"context"
	"errors"
	"github.com/AriSu2904/go-auth/internal/config"
	"log"
	"log/slog"
	"strings"

	"github.com/AriSu2904/go-auth/internal/dto"
	"github.com/AriSu2904/go-auth/internal/models"
	"github.com/AriSu2904/go-auth/internal/repository"
	"github.com/AriSu2904/go-auth/internal/utils"
)

var (
	ErrUserExists         = errors.New("user with this email or persona already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService interface {
	SignUp(ctx context.Context, input *dto.RegisterUserInput) (*models.User, error)
	SignIn(ctx context.Context, input *dto.LoginUserInput, additionalHeader *dto.AdditionalHeader) (*models.TokenInfo, error)
}

type authService struct {
	userRepository repository.UserRepository
	logger         *slog.Logger
	config         *config.Config
}

func NewAuthService(userRepo repository.UserRepository, log *slog.Logger, conf *config.Config) AuthService {
	return &authService{userRepository: userRepo, logger: log, config: conf}
}

func (s *authService) SignUp(ctx context.Context,
	input *dto.RegisterUserInput) (*models.User, error) {
	s.logger.Info("executing signup request", "layer", "authService")

	email := input.Email
	existingUser, err := s.userRepository.FindByEmail(ctx, &email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		slog.Error("Error checking existing user by email:", err)

		return nil, err
	}

	if existingUser != nil {
		slog.Warn("Error user already exist", err)

		return nil, ErrUserExists
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:              input.Email,
		Persona:            input.Persona,
		Password:           hashedPassword,
		Role:               "USER",
		Status:             "ACTIVE",
		IsVerified:         false,
		GoogleSynchronized: false,
	}

	err = s.userRepository.Create(ctx, newUser)
	if err != nil && err.Error() != "sql: no rows in result set" {
		slog.Error("Error creating new user:", err)

		return nil, err
	}

	return newUser, nil
}

func (s *authService) SignIn(ctx context.Context,
	input *dto.LoginUserInput,
	additionalHeader *dto.AdditionalHeader,
) (*models.TokenInfo, error) {
	s.logger.Info("executing login request", "layer", "authService")

	isUsingEmail := strings.Contains(input.UniqueId, "@")
	var user *models.User

	if isUsingEmail {
		existUser, err := s.userRepository.FindByEmail(ctx, &input.UniqueId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return nil, ErrUserNotFound
			} else {
				log.Println("Error occurred when find user by email:", err)
				return nil, err
			}
		}

		user = existUser
	} else {
		existUser, err := s.userRepository.FindByPersona(ctx, &input.UniqueId)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return nil, ErrUserNotFound
			} else {
				log.Println("Error occurred when find user by persona:", err)
				return nil, err
			}
		}

		user = existUser
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		return nil, ErrInvalidCredentials
	}

	tokenInfo, err := utils.GenerateTokenJwt(user, s.config.PrivateKey, &utils.ExpiryTime{
		AccessExpiry:  s.config.JwtAccessTokenExpiry,
		RefreshExpiry: s.config.JwtRefreshTokenExpiry,
	})
	if err != nil {
		return nil, err
	}

	return tokenInfo, nil
}
