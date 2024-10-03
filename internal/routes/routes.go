package routes

import (
	"NewsAggregator/internal/articles"
	"NewsAggregator/internal/users"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Articles
	mux.HandleFunc("/articles", articles.GetAllArticlesHandler)        // GET all articles
	mux.HandleFunc("/articles/{id}", articles.GetArticleHandler)       // GET article by id
	mux.HandleFunc("/articles/search", articles.SearchArticlesHandler) // GET articles by search

	mux.HandleFunc("/articles/like/{id}", articles.LikeArticleHandler) // POST like/dislike article

	mux.HandleFunc("/articles/comment/{id}", articles.CommentArticleHandler)         // POST comment article
	mux.HandleFunc("/articles/comments/{id}", articles.GetCommentsForArticleHandler) // GET comments

	mux.HandleFunc("/articles/read-later/{id}", articles.ReadLaterArticleHandler) // POST read later/remove from read later article
	mux.HandleFunc("/articles/read-later", articles.GetReadLaterArticlesHandler)  // GET read later articles

	mux.HandleFunc("/articles/category/{name}", articles.GetArticlesByCategoryHandler) // GET articles by category
	mux.HandleFunc("/articles/category", articles.GetCategoriesHandler)                // GET categories

	// Users
	mux.HandleFunc("/user/register", users.RegisterUserHandler)
	mux.HandleFunc("/user/login", users.LoginUserHandler)

	return mux
}
