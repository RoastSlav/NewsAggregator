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

	body, err := io.ReadAll(get.Body)
	Util.CheckErrorAndLog(err, "Failed to read response body")

	var newsAPIResponse NewsAPIResponse
	err = json.Unmarshal(body, &newsAPIResponse)
	Util.CheckErrorAndLog(err, "Failed to unmarshal response body")

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

	articles, err := GetAllArticles()
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get all articles", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode articles", http.StatusInternalServerError)
}
