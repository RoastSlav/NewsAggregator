package routes

import (
	"NewsAggregator/internal/articles"
	"NewsAggregator/internal/categories"
	"NewsAggregator/internal/users"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Articles
	mux.HandleFunc("/articles", articles.GetAllArticlesHandler)
	mux.HandleFunc("/articles/{id}", articles.GetArticleHandler)
	mux.HandleFunc("/articles/search", articles.SearchArticlesHandler)
	mux.HandleFunc("/articles/like/{id}", articles.LikeArticleHandler)
	mux.HandleFunc("/articles/comment/{id}", articles.CommentArticleHandler)
	mux.HandleFunc("/articles/read-later/{id}", articles.ReadLaterArticleHandler)

	// Categories
	mux.HandleFunc("/category/{name}", categories.GetArticlesByCategoryHandler)
	mux.HandleFunc("/category", categories.GetCategoriesHandler)

	// Users
	mux.HandleFunc("/user/register", users.RegisterUserHandler)
	mux.HandleFunc("/user/login", users.LoginUserHandler)

	return mux
}
