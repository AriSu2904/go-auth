package models

import (
	"github.com/AriSu2904/go-auth/internal/types"
	"time"
)

type UserRole string
type UserStatus string

const (
	RoleUser  UserRole = "USER"
	RoleAdmin UserRole = "ADMIN"

	StatusActive   UserStatus = "ACTIVE"
	StatusInactive UserStatus = "INACTIVE"
	StatusBanned   UserStatus = "BANNED"
)

type User struct {
	ID                 string           `json:"id" db:"id"`
	FirstName          types.NullString `json:"firstName,omitempty" db:"first_name"`
	LastName           types.NullString `json:"lastName,omitempty" db:"last_name"`
	Dob                types.NullString `json:"dob,omitempty" db:"dob"`
	Email              string           `json:"email" db:"email"`
	Persona            string           `json:"persona" db:"persona"`
	Password           string           `json:"-" db:"password"`
	Role               UserRole         `json:"role" db:"role"`
	IsVerified         bool             `json:"isVerified" db:"is_verified"`
	GoogleSynchronized bool             `json:"googleSynchronized" db:"google_synchronized"`
	Status             UserStatus       `json:"status" db:"status"`
	CreatedAt          time.Time        `json:"createdAt" db:"created_at"`
	ModifiedAt         time.Time        `json:"modifiedAt" db:"modified_at"`
}
