package rapyd

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rAlexander89/rapyd-banking/countries"
)

func TestGetCountries(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	sk := os.Getenv("SECRET_KEY")
	ak := os.Getenv("ACCESS_KEY")

	cli := http.DefaultClient
	url, _ := url.Parse(rapydSandbox)
	signer := NewSigner([]byte(ak), []byte(sk))
	client := NewClient(signer, cli, url)

	var want *countries.GetCountriesResponse
	response, err := client.GetCountries()

	if reflect.TypeOf(response) != reflect.TypeOf(want) {
		t.Error("Get Countries Error: ", err)
	}
}
