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

func CheckHttpMethodAndSendHttpResponse(r *http.Request, w http.ResponseWriter, method string, message string, responseCode int) bool {
	if r.Method != method {
		http.Error(w, message, responseCode)
		log.Println(message)
		return true
	}
	return false
}

func CheckEmptyAndSendHttpResponse(value string, w http.ResponseWriter, message string, responseCode int) bool {
	if value == "" {
		http.Error(w, message, responseCode)
		log.Println(message)
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

func CheckErrorAndFatal(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}
