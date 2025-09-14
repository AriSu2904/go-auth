package service

import (
	"context"
	"errors"

	"github.com/AriSu2904/go-auth/internal/dto"
	"github.com/AriSu2904/go-auth/internal/models"
	"github.com/AriSu2904/go-auth/internal/repository"
	"github.com/AriSu2904/go-auth/internal/utils"
)

var (
	ErrUserExists = errors.New("user with this email or persona already exists")
)

type AuthService interface {
	SignUp(ctx context.Context, input *dto.RegisterUserInput) (*models.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepository: userRepo}
}

func (s *authService) SignUp(ctx context.Context, input *dto.RegisterUserInput) (*models.User, error) {
	existingUser, err := s.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
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
		GoogleSynchronized: false,
		IsVerified:         false,
		Status:             "ACTIVE",
		Role:               "USER",
	}

	err = s.userRepository.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
