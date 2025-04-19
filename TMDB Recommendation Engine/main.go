package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const base_url = "https://api.themoviedb.org/3"

func getURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to get env variables")
		return ""
	}
	API_KEY := os.Getenv("API_KEY")
	return API_KEY
}

func main() {
	apikey := getURI()
	fmt.Println(apikey)

}
