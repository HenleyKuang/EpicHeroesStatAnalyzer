package main

import (
	"fmt"
	"log"
	"main/api"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Printf("Starting service in port: %s\n", port)
	api.HandleRequests(port)
}
