package Api

import (
	"NewsAggregator/internal/articles"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := articles.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to get all articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		log.Printf("Failed to encode articles: %v", err)
		http.Error(w, "Failed to encode articles", http.StatusInternalServerError)
	}
}
