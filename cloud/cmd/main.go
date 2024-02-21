package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/pkulik0/stredono/cloud"
	"log"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("PORT not set. Using default (%s).", defaultPort)
		port = defaultPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalln("Failed to serve Cloud Functions: ")
	}
}
