package main

import (
	"fmt"
	"github.com/ShiraazMoollatjie/devtogo"
	"log"
	"os"
)

func main() {
	apiToken := os.Getenv("API_KEY")
	if apiToken == "" {
		log.Fatalf("Set the API_KEY")
	}
	dtgc := devtogo.NewClient(devtogo.WithApiKey(apiToken))

	al, err := dtgc.Articles(devtogo.Arguments{
		"tag": "go",
	})
	if err != nil {
		log.Fatalf("Cannot retrieve articles: %+v", err)
	}

	for _, a := range al {
		fmt.Printf("Title: %s, Url: %s\n", a.Title, a.URL)
	}
}
