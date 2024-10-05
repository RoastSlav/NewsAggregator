package users

import (
	"NewsAggregator/internal/database"
)

func insertUser(user *User) error {
	_, err := database.DB.NamedExec("INSERT INTO users (username, email, password_hash) VALUES (:username, :email, :password_hash)", user)
	return err
}

func getUserByEmail(email string) (User, error) {
	var user User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	return user, err
}

func getUserIdBySessionToken(sessionToken string) (int, error) {
	var userId int
	err := database.DB.Get(&userId, "SELECT sessions.user_id FROM sessions WHERE session_token = ?", sessionToken)
	return userId, err
}
