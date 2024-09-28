package users

import "NewsAggregator/database"

func InsertUser(user *User) error {
	_, err := database.DB.NamedExec("INSERT INTO users (username, email, password_hash) VALUES (:username, :email, :password_hash)", user)
	return err
}
