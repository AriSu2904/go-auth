package models

import (
	"github.com/lib/pq"
	"time"
)

type Session struct {
	ID           string         `json:"id" db:"id"`
	UserID       string         `json:"userId" db:"user_id"`
	DeviceID     string         `json:"deviceId,omitempty" db:"device_id"`
	RefreshToken string         `json:"-" db:"refresh_token"`
	IPAddress    pq.StringArray `json:"-" db:"ip_address"`
	UserAgent    string         `json:"userAgent,omitempty" db:"user_agent"`
	CreatedAt    time.Time      `json:"createdAt" db:"created_at"`
	ExpiredAt    time.Time      `json:"expiredAt" db:"expired_at"`
}
