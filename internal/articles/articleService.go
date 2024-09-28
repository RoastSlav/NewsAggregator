package articles

import (
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
	if err != nil {
		log.Fatalf("Failed to fetch article: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(get.Body)

	body, err := io.ReadAll(get.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var newsAPIResponse NewsAPIResponse
	err = json.Unmarshal(body, &newsAPIResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if newsAPIResponse.Status != "ok" {
		log.Fatalf("Failed to fetch articles from News API: %s", newsAPIResponse.Message)
	}

	for _, article := range newsAPIResponse.Articles {
		article.CreatedAt = time.Now()
		err = insertArticle(&article)
		if err != nil {
			log.Fatalf("Failed to insert article: %v", err)
		}
	}

	log.Printf("Fetched %d articles from News API", newsAPIResponse.TotalResults)
}

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := GetAllArticles()
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
