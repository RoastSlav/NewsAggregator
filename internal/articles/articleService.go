package articles

import (
	"NewsAggregator/internal/users"
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

	categoryDB, err := getCategoryByName(categoryEnv)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Fatalf("Failed to get category by name: %v", err)
	}

	if categoryDB.Name == "" {
		category := Category{
			Name: categoryEnv,
		}
		err = insertCategory(&category)
		Util.CheckErrorAndLog(err, "Failed to insert category")
		categoryDB, err = getCategoryByName(categoryEnv)
	}

	for _, article := range newsAPIResponse.Articles {
		articleDB, err := getArticlesByTitleAndAuthor(article.Title, article.Author)
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
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

	articleID := r.PathValue("id")

	id, _ := strconv.ParseInt(articleID, 0, 64)
	article, err := getArticleById(int(id))
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get article", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(article)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode article", http.StatusInternalServerError)
}

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var everyArticleRequest PagedRequest
	err := json.NewDecoder(r.Body).Decode(&everyArticleRequest)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	if everyArticleRequest.Page < 1 {
		http.Error(w, "Page must be greater than 0", http.StatusBadRequest)
	}

	articles, err := getAllArticles(everyArticleRequest.Page, everyArticleRequest.Limit)
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
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

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

	articles, err := searchArticles(searchArticleRequest)
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

func GetArticlesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

	categoryName := r.PathValue("name")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to close request body", http.StatusInternalServerError)
	}(r.Body)

	var categoryArticlesRequest PagedRequest
	err := json.NewDecoder(r.Body).Decode(&categoryArticlesRequest)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)

	articles, err := getArticlesByCategoryName(categoryName, categoryArticlesRequest.Limit, categoryArticlesRequest.Page)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get category by name", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode articles", http.StatusInternalServerError)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

	categories, err := getCategories()
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to get categories", http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to encode categories", http.StatusInternalServerError)
}

func LikeArticleHandler(w http.ResponseWriter, r *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodPost, "Invalid request method", http.StatusMethodNotAllowed)

	articleID := r.PathValue("id")
	id, _ := strconv.ParseInt(articleID, 0, 64)

	sessionToken := r.Header.Get("Authorization")
	Util.CheckEmptyAndSendHttpResponse(sessionToken, w, "Authorization header is required", http.StatusUnauthorized)

	userId := users.GetUserIdFromSessionToken(sessionToken)
	liked, err := checkIfUserLikedArticle(int(id), userId)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to check if user liked article", http.StatusInternalServerError)

	if liked {
		err := deleteLikeFromArticle(int(id), userId)
		Util.CheckErrorAndSendHttpResponse(err, w, "Failed to remove like from article", http.StatusInternalServerError)
		return
	}
	err = addLikeToArticle(int(id), userId)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to add like to article", http.StatusInternalServerError)
}

func CommentArticleHandler(w http.ResponseWriter, r *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(r, w, http.MethodPost, "Invalid request method", http.StatusMethodNotAllowed)

	articleID := r.PathValue("id")
	id, _ := strconv.ParseInt(articleID, 0, 64)

	sessionToken := r.Header.Get("Authorization")
	Util.CheckEmptyAndSendHttpResponse(sessionToken, w, "Authorization header is required", http.StatusUnauthorized)

	userId := users.GetUserIdFromSessionToken(sessionToken)

	var comment CommentRequest
	err := json.NewDecoder(r.Body).Decode(&comment)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to decode request body", http.StatusBadRequest)
	Util.CheckEmptyAndSendHttpResponse(comment.Content, w, "CommentRequest content is required", http.StatusBadRequest)

	err = addCommentToArticle(int(id), userId, comment.Content)
	Util.CheckErrorAndSendHttpResponse(err, w, "Failed to insert comment", http.StatusInternalServerError)
}

func ReadLaterArticleHandler(writer http.ResponseWriter, request *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(request, writer, http.MethodPost, "Invalid request method", http.StatusMethodNotAllowed)

	articleID := request.PathValue("id")
	id, _ := strconv.ParseInt(articleID, 0, 64)

	sessionToken := request.Header.Get("Authorization")
	Util.CheckEmptyAndSendHttpResponse(sessionToken, writer, "Authorization header is required", http.StatusUnauthorized)

	userID := users.GetUserIdFromSessionToken(sessionToken)

	inReadLater, err := isArticleInReadLater(userID, int(id))
	Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to check if article is in read later", http.StatusInternalServerError)

	if inReadLater {
		err := removeArticleFromReadLater(userID, int(id))
		Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to remove article from read later", http.StatusInternalServerError)
		return
	}

	err = addArticleToReadLater(userID, int(id))
	Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to add article to read later", http.StatusInternalServerError)
}

func GetReadLaterArticlesHandler(writer http.ResponseWriter, request *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(request, writer, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

	sessionToken := request.Header.Get("Authorization")
	Util.CheckEmptyAndSendHttpResponse(sessionToken, writer, "Authorization header is required", http.StatusUnauthorized)

	userID := users.GetUserIdFromSessionToken(sessionToken)

	articles, err := getReadLaterArticles(userID)
	Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to get read later articles", http.StatusInternalServerError)

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(articles)
	Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to encode articles", http.StatusInternalServerError)
}

func GetCommentsForArticleHandler(writer http.ResponseWriter, request *http.Request) {
	Util.CheckHttpMethodAndSendHttpResponse(request, writer, http.MethodGet, "Invalid request method", http.StatusMethodNotAllowed)

	articleID := request.PathValue("id")
	id, _ := strconv.ParseInt(articleID, 0, 64)

	comments, err := getCommentsForArticle(int(id))
	Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to get comments for article", http.StatusInternalServerError)

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(comments)
	Util.CheckErrorAndSendHttpResponse(err, writer, "Failed to encode comments", http.StatusInternalServerError)
}
