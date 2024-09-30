package articles

import (
	"NewsAggregator/internal/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func FetchArticlesFromNewsAPI(topic string) {
	apiKey := os.Getenv("NEWS_API_KEY")
	newsEverythingEndpointUrl := os.Getenv("NEWS_API_EVERYTHING_ENDPOINT_URL")

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

	for _, article := range newsAPIResponse.Articles {
		article.CreatedAt = time.Now()
		err = insertArticle(&article)
		Util.CheckErrorAndLog(err, "Failed to insert article")
	}

	log.Printf("Fetched %d articles from News API", newsAPIResponse.TotalResults)
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
