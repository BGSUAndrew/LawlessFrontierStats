package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found")
	}

	apiKey := os.Getenv("Lawless_API_Key")
	if apiKey == "" {
		log.Fatal("Lawless_API_Key not set")
	}

	url := "https://www.bungie.net/Platform/Destiny2/1/Account/4611686018430407175/Character/2305843009263048442/Stats/Activities/?mode=93"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
