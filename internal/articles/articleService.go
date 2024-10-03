package articles

import (
	"NewsAggregator/internal/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func FetchArticlesFromNewsAPI(topic string) {
	apiKey := os.Getenv("NEWS_API_KEY")
	newsEverythingEndpointUrl := os.Getenv("NEWS_API_EVERYTHING_ENDPOINT_URL")
	categoryEnv := os.Getenv("NEWS_API_TOPIC")

	url := fmt.Sprintf(newsEverythingEndpointUrl+"q=%s&apiKey=%s", topic, apiKey)

	get, err := http.Get(url)
	Util.CheckErrorAndLog(err, "Failed to fetch articles from News API")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndLog(err, "Failed to close response body")
	}(get.Body)

	var newsAPIResponse NewsAPIResponse
	Util.ParseBodyFromJson(get, &newsAPIResponse)

	if newsAPIResponse.Status != "ok" {
		log.Fatalf("Failed to fetch articles from News API: %s", newsAPIResponse.Message)
	}

	categoryDB, err := GetCategoryByName(categoryEnv)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Fatalf("Failed to get category by name: %v", err)
	}

	if categoryDB.Name == "" {
		category := Category{
			Name: categoryEnv,
		}
		err = InsertCategory(&category)
		Util.CheckErrorAndLog(err, "Failed to insert category")
		categoryDB, err = GetCategoryByName(categoryEnv)
	}

	for _, article := range newsAPIResponse.Articles {
		articleDB, err := GetArticlesByTitleAndAuthor(article.Title, article.Author)
		Util.CheckErrorAndLog(err, "Failed to get article by title and author")

		if len(articleDB) > 0 {
			continue
		}

		// Remove articles with title "[Removed]" as they are not useful. The API sometimes returns articles with this title.
		if article.Title == "[Removed]" {
			continue
		}

		article.CreatedAt = time.Now()
		err = insertArticle(&article, categoryDB.ID)
		Util.CheckErrorAndLog(err, "Failed to insert article")
	}

	log.Printf("Fetched %d articles from News API", newsAPIResponse.TotalResults)
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	articleID := r.PathValue("id")

	id, _ := strconv.ParseInt(articleID, 0, 64)
	article, err := GetArticleById(int(id))
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get article", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(article)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode article", http.StatusInternalServerError)
}

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var everyArticleRequest EveryArticleRequest
	err := json.NewDecoder(r.Body).Decode(&everyArticleRequest)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	if everyArticleRequest.Page < 1 {
		http.Error(w, "Page must be greater than 0", http.StatusBadRequest)
	}

	articles, err := GetAllArticles(everyArticleRequest.Page, everyArticleRequest.Limit)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get all articles", http.StatusInternalServerError)

	response := EveryArticleResponse{
		Articles:     articles,
		TotalResults: len(articles),
		Page:         everyArticleRequest.Page,
		Limit:        everyArticleRequest.Limit,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode articles", http.StatusInternalServerError)
}

func SearchArticlesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var searchArticleRequest SearchArticleRequest
	err := json.NewDecoder(r.Body).Decode(&searchArticleRequest)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	if searchArticleRequest.PublishedFrom.After(searchArticleRequest.PublishedTo) {
		http.Error(w, "PublishedFrom date cannot be after PublishedTo date", http.StatusBadRequest)
	}

	if searchArticleRequest.Page < 1 {
		http.Error(w, "Page must be greater than 0", http.StatusBadRequest)
	}

	articles, err := SearchArticles(searchArticleRequest)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to search articles", http.StatusInternalServerError)

	response := SearchArticleResponse{
		Articles:     articles,
		TotalResults: len(articles),
		Page:         searchArticleRequest.Page,
		Limit:        searchArticleRequest.Limit,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode articles", http.StatusInternalServerError)
}
