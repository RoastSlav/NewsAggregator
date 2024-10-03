package categories

import (
	"NewsAggregator/internal/articles"
	Util "NewsAggregator/internal/util"
	"encoding/json"
	"io"
	"net/http"
)

func GetArticlesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	categoryName := r.PathValue("name")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var categoryArticlesRequest PagedRequest
	err := json.NewDecoder(r.Body).Decode(&categoryArticlesRequest)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	articles, err := articles.GetArticlesByCategoryName(categoryName, categoryArticlesRequest.Limit, categoryArticlesRequest.Page)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get category by name", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode articles", http.StatusInternalServerError)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	categories, err := GetCategories()
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get categories", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode categories", http.StatusInternalServerError)
}
