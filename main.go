package main

import (
	"NewsAggregator/database"
	"NewsAggregator/internal/articles"
	"fmt"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	articles.FetchArticlesFromNewsAPI()

	allArticles, err := articles.GetAllArticles()
	if err != nil {
		log.Fatalf("Failed to get all articles: %v", err)
	}

	for _, article := range allArticles {
		fmt.Printf("Title: %s Published at: %v\n", article.Title, article.PublishedAt)
	}
}
