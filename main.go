package main

import (
	"NewsAggregator/database"
	"NewsAggregator/internal"
	"NewsAggregator/internal/articles"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	articles.FetchArticlesFromNewsAPI()

	err = Api.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
