package users

import (
	"NewsAggregator/internal/database"
)

func InsertUser(user *User) error {
	_, err := database.DB.NamedExec("INSERT INTO users (username, email, password_hash) VALUES (:username, :email, :password_hash)", user)
	return err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	return user, err
}
