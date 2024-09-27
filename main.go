package main

import (
	"NewsAggregator/database"
	"log"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

}
