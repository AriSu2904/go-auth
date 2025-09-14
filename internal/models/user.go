package models

import "time"

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
	ID                 string     `json:"id" db:"id"`
	FirstName          string     `json:"firstName,omitempty" db:"first_name"`
	LastName           string     `json:"lastName,omitempty" db:"last_name"`
	Email              string     `json:"email" db:"email"`
	Persona            string     `json:"persona" db:"persona"`
	Password           string     `json:"-" db:"password"`
	Role               UserRole   `json:"role" db:"role"`
	IsVerified         bool       `json:"isVerified" db:"is_verified"`
	GoogleSynchronized bool       `json:"googleSynchronized" db:"google_synchronized"`
	Status             UserStatus `json:"status" db:"status"`
	CreatedAt          time.Time  `json:"createdAt" db:"created_at"`
	ModifiedAt         time.Time  `json:"modifiedAt" db:"modified_at"`
}
