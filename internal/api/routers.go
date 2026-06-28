package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	// api_key := os.Getenv("API_KEY")
	// fmt.Println(api_key)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	url := "https://epic.gsfc.nasa.gov/api/natural"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
