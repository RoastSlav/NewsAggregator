package users

import "time"

type User struct {
	ID           int       `db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"passwordHash" db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserRegistration struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	SessionToken string    `db:"session"`
	CreatedAt    time.Time `db:"created_at"`
}
