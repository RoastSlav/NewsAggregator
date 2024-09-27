package Api

import (
	"log"
	"net/http"
)

func registerHandlers() {
	http.HandleFunc("/articles", GetAllArticlesHandler)
}

func StartServer() error {
	registerHandlers()
	log.Println("Handler registered")

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
