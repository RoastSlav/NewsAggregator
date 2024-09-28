package users

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"time"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, "Failed to close request body", http.StatusInternalServerError)
		}
	}(r.Body)

	var user UserRegistration
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	userDB := User{
		Username:     user.Username,
		PasswordHash: string(hashedPassword),
		Email:        user.Email,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := InsertUser(&userDB); err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	log.Fatalln("User registered")
}
