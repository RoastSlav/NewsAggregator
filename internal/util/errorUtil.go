package Util

import (
	"log"
	"net/http"
)

func CheckErrorAndSendHttpResponse(err error, w http.ResponseWriter, message string, responseCode int) bool {
	if err != nil {
		http.Error(w, message, responseCode)
		log.Println(message, err)
		return true
	}
	return false
}

func CheckErrorAndLog(err error, message string) bool {
	if err != nil {
		log.Printf("%s: %v\n", message, err)
		return true
	}
	return false
}
