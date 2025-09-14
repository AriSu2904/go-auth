package models

import (
	"time"
)

type Session struct {
	ID           string    `json:"id" db:"id"`
	UserID       string    `json:"userId" db:"user_id"`
	DeviceID     string    `json:"deviceId" db:"device_id"`
	DeviceInfo   string    `json:"deviceInfo,omitempty" db:"device_info"`
	RefreshToken string    `json:"-" db:"refresh_token"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	ExpiredAt    time.Time `json:"expiredAt" db:"expired_at"`
}
