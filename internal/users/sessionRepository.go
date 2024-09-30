package users

import (
	"NewsAggregator/internal/database"
	"time"
)

func InsertSessionToken(token string, userId int) error {
	_, err := database.DB.Exec("INSERT INTO sessions (session_token, user_id, created_at) VALUES (?, ?, ?)", token, userId, time.Now())
	return err
}
