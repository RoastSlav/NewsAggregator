package routes

import (
	"NewsAggregator/internal/articles"
	"NewsAggregator/internal/users"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/articles", articles.GetAllArticlesHandler)
	mux.HandleFunc("/articles/search", articles.SearchArticlesHandler)

	mux.HandleFunc("/user/register", users.RegisterUserHandler)
	mux.HandleFunc("/user/login", users.LoginUserHandler)

	return mux
}
