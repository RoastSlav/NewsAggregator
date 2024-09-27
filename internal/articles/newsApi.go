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

func FetchArticlesFromNewsAPI() {
	apiKey := os.Getenv("NEWS_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=technology&apiKey=%s", apiKey)

	get, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch article: %v", err)
	}

	defer get.Body.Close()

	body, err := io.ReadAll(get.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var newsAPIResponse NewsAPIResponse
	err = json.Unmarshal(body, &newsAPIResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
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
