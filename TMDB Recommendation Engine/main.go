package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const base_url = "https://api.themoviedb.org/3/"

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

	resp, err := http.Get(base_url + apikey)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
	}

	// if err != json.NewDecoder(resp.Body) {
	// 	continue
	// }

}
