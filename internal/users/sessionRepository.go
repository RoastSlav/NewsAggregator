package users

import (
	"NewsAggregator/internal/database"
)

func insertSessionToken(token string, userId int) error {
	_, err := database.DB.Exec("INSERT INTO sessions (session_token, user_id) VALUES (?, ?)", token, userId)
	return err
}

func getSessionByToken(token string) (Session, error) {
	var session Session
	err := database.DB.Get(&session, "SELECT * FROM sessions WHERE session_token = ?", token)
	return session, err
}
