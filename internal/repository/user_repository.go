package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/AriSu2904/go-auth/internal/models"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByPersona(ctx context.Context, persona *string) (*models.User, error)
	FindById(ctx context.Context, id string) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (userRepository *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, persona, password, role, status, is_verified, google_synchronized)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := userRepository.DB.QueryRowContext(ctx, query,
		user.Email, user.Persona, user.Password, user.Role, user.Status, user.IsVerified, user.GoogleSynchronized,
	).Scan(&user.ID, &user.CreatedAt, &user.ModifiedAt)

	return err
}

func (userRepository *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, persona, password, role, is_verified, google_synchronized, status, created_at, modified_at
			  FROM users WHERE email = $1`

	row := userRepository.DB.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Persona,
		&user.Password,
		&user.Role,
		&user.IsVerified,
		&user.GoogleSynchronized,
		&user.Status,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) FindByPersona(ctx context.Context, persona *string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, persona, password, role, is_verified, google_synchronized, status, created_at, modified_at
			  FROM users WHERE persona = $1`

	row := userRepository.DB.QueryRowContext(ctx, query, persona)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Persona,
		&user.Password,
		&user.Role,
		&user.IsVerified,
		&user.GoogleSynchronized,
		&user.Status,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepository *userRepository) FindById(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, persona, password, role, is_verified, google_synchronized, status, created_at, modified_at
			  FROM users WHERE id = $1`

	row := userRepository.DB.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Persona,
		&user.Password,
		&user.Role,
		&user.IsVerified,
		&user.GoogleSynchronized,
		&user.Status,
		&user.CreatedAt,
		&user.ModifiedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
