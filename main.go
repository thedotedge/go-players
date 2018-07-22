package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	port := "5000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := http.ListenAndServe(":" + port, server); err != nil {
		log.Fatalf("could not listen on port %s %v", port, err)
	}
}
