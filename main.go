package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	sk := os.Getenv("SECRET_KEY")
	ak := os.Getenv("ACCESS_KEY")

	cli := http.DefaultClient
	url, _ := url.Parse("https://sandboxapi.rapyd.net")
	signer := NewSigner([]byte(ak), []byte(sk))
	client := NewClient(signer, cli, url)
	response, _ := client.GetCountries()

	fmt.Println("final response: ", response)
}
