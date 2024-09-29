package Api

import (
	"NewsAggregator/internal/articles"
	"NewsAggregator/internal/users"
	Util "NewsAggregator/internal/util"
	"log"
	"net/http"
)

func registerHandlers() {
	http.HandleFunc("/articles", articles.GetAllArticlesHandler)
	http.HandleFunc("/articles/search", articles.SearchArticlesHandler)

	http.HandleFunc("/user/register", users.RegisterUserHandler)
	http.HandleFunc("/user/login", users.LoginUserHandler)
}

func StartServer() error {
	registerHandlers()
	log.Println("Handler registered")

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	errored := Util.CheckErrorAndLog(err, "Failed to start server")
	if errored {
		return err
	}
	return nil
}
