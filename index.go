package main

import (
	"log"
	"net/http"
	"os"

	"index/bungie"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("Lawless_API_Key")
	if apiKey == "" {
		log.Fatal("Key not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/activities", func(w http.ResponseWriter, r *http.Request) {
		bungie.LawlessData(w, r, apiKey)
	})

	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
