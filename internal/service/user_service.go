package service

import (
	"context"
	"errors"
	"github.com/AriSu2904/go-auth/internal/models"
	"github.com/AriSu2904/go-auth/internal/repository"
	"log"
	"log/slog"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	FindByPersona(ctx context.Context, persona *string) (*models.User, error)
	FindByEmail(ctx context.Context, email *string) (*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
	logger         *slog.Logger
}

func NewUserService(userRepo repository.UserRepository, log *slog.Logger) UserService {
	return &userService{userRepository: userRepo, logger: log}
}

func (s *userService) FindByPersona(ctx context.Context, persona *string) (*models.User, error) {
	s.logger.Info("executing find user by persona", "layer", "userService")

	user, err := s.userRepository.FindByPersona(ctx, persona)

	if err != nil {
		log.Println("Error occurred when find user by persona:", err)
		if err.Error() == "sql: no rows in result set" {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *userService) FindByEmail(ctx context.Context, email *string) (*models.User, error) {
	s.logger.Info("executing find user by email", "layer", "userService")

	user, err := s.userRepository.FindByEmail(ctx, email)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
	}

	return user, nil
}
