package main

import (
	"NewsAggregator/internal/articles"
	"NewsAggregator/internal/database"
	"NewsAggregator/internal/routes"
	Util "NewsAggregator/internal/util"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv()

	err := database.Connect()
	Util.CheckErrorAndFatal(err, "Failed to connect to database")

	topic := os.Getenv("NEWS_API_TOPIC")
	articles.FetchArticlesFromNewsAPI(topic)

	httpHandler := routes.NewRouter()
	port := 8080

	err = initCronJobs()
	Util.CheckErrorAndFatal(err, "Failed to initialize cron jobs")

	log.Printf("Server is listening on http://localhost:%v\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), httpHandler)
	Util.CheckErrorAndFatal(err, "Failed to start server")
}

func loadEnv() {
	err := godotenv.Load("config.env")
	Util.CheckErrorAndFatal(err, "Failed to load environment variables")
}

func initCronJobs() error {
	c := cron.New()

	_, err := c.AddJob("0 */1 * * *", articles.FetchJob{})
	if err != nil {
		return err
	}

	c.Start()
	return nil
}
