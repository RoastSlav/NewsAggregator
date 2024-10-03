package users

import (
	"NewsAggregator/internal/util"
	"crypto/rand"
	"encoding/base64"
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
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var user UserRegistration
	err := json.NewDecoder(r.Body).Decode(&user)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to hash password", http.StatusInternalServerError)

	userDB := User{
		Username:     user.Username,
		PasswordHash: string(hashedPassword),
		Email:        user.Email,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = InsertUser(&userDB)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to insert user", http.StatusInternalServerError)

	log.Printf("User registered. Username:%s, Email:%s \n", userDB.Username, userDB.Email)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var userLogin UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	user, err := GetUserByEmail(userLogin.Email)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get user", http.StatusInternalServerError)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userLogin.Password))
	Util.CheckErrorAndSendHttpResponse(err, w, "Invalid password", http.StatusUnauthorized)

	log.Println("User logged in")

	sessionToken := generateSessionToken()
	err = InsertSessionToken(sessionToken, user.ID)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to insert session token", http.StatusInternalServerError)

	w.Header().Set("Authorization", sessionToken)
}

func generateSessionToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	Util.CheckErrorAndLog(err, "Failed to generate session token")

	return base64.StdEncoding.EncodeToString(bytes)
}

func GetUserIdFromSessionToken(sessionToken string) int {
	id, err := getUserIdBySessionToken(sessionToken)
	Util.CheckErrorAndLog(err, "Failed to get user id from session token")

	return id
}
