package users

import "time"

type Type int

const (
	TypeInvalid         = Type(iota)
	TypeSecurityOfficer = Type(iota)
	TypeUser            = Type(iota)
)

type User struct {
	ID                   int64      `json:"id" db:"id"`
	Type                 Type       `json:"type" db:"type"`
	Username             string     `json:"username" db:"username"`
	HashedPassword       string     `json:"-" db:"hashed_password"`
	LastSucceededLoginAt *time.Time `json:"last_succeeded_login_at" db:"last_succeeded_login_at"`
	LastFailedLoginAt    *time.Time `json:"last_failed_login_at" db:"last_failed_login_at"`
	PasswordExpiredAt    time.Time  `json:"password_expired_at" db:"password_expired_at"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`
}

const DefaultUser = "default"
const DefaultSecurityOfficer = "default_security_officer"
