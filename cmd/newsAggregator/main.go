package main

import (
	"NewsAggregator/internal/articles"
	"NewsAggregator/internal/database"
	"NewsAggregator/internal/routes"
	Util "NewsAggregator/internal/util"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	articles.FetchArticlesFromNewsAPI("technology")

	httpHandler := routes.NewRouter()
	port := 8080

	log.Printf("Server is listening on http://localhost:%v\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), httpHandler)

	Util.CheckErrorAndLog(err, "Failed to start server")
}
