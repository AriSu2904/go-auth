package repository

import (
	"context"
	"database/sql"
	"github.com/AriSu2904/go-auth/internal/models"
	"log/slog"
	"time"
)

type SessionRepository interface {
	Create(ctx context.Context, session *models.Session) error
	FindByDeviceId(ctx context.Context, deviceId *string) (*models.Session, error)
}

type sessionRepository struct {
	Db     *sql.DB
	logger *slog.Logger
}

func NewSessionRepository(db *sql.DB, log *slog.Logger) SessionRepository {
	return &sessionRepository{Db: db, logger: log}
}

func (s *sessionRepository) FindByDeviceId(ctx context.Context, deviceId *string) (*models.Session, error) {
	s.logger.Info("Finding session by device id", "layer", "sessionRepository", "deviceId", *deviceId)

	query := `SELECT id, user_id, device_id, device_info, refresh_token, expired_at, created_at, modified_at
			  FROM sessions WHERE device_id = $1`

	row := s.Db.QueryRowContext(ctx, query, *deviceId)

	var session models.Session
	err := row.Scan(
		&session.ID,
		&session.UserID,
		&session.DeviceID,
		&session.DeviceInfo,
		&session.RefreshToken,
		&session.ExpiredAt,
		&session.CreatedAt,
		&session.ModifiedAt,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *sessionRepository) Create(ctx context.Context, session *models.Session) error {
	s.logger.Info("Creating new session", "layer", "sessionRepository", "deviceId", session.DeviceID)

	query := `INSERT INTO sessions (user_id, device_id, device_info, refresh_token, expired_at)
			  VALUES ($1, $2, $3, $4, $5)`

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err := s.Db.QueryRowContext(ctx, query,
		session.UserID, session.DeviceID, session.RefreshToken, session.ExpiredAt,
	).Scan(&session.ID, &session.CreatedAt, &session.ModifiedAt)

	return err
}
