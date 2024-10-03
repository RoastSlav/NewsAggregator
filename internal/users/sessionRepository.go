package users

import (
	"NewsAggregator/internal/database"
)

func InsertSessionToken(token string, userId int) error {
	_, err := database.DB.Exec("INSERT INTO sessions (session_token, user_id) VALUES (?, ?)", token, userId)
	return err
}
