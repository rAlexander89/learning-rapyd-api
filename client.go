package rapyd

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/rAlexander89/rapyd-banking/countries"
)

type RapydClient struct {
	Signer
	*http.Client
	url *url.URL
}

type Client interface {
	Sign(path string) ([]byte, error)
	GetCountries() (*countries.GetCountriesResponse, error)
	Resolve(path string) string
}

func NewClient(signer Signer, cli *http.Client, url *url.URL) Client {
	return &RapydClient{
		Signer: signer,
		Client: cli,
		url:    url,
	}
}

func (c *RapydClient) Resolve(path string) string {
	endpoint, err := url.Parse(path)
	if err != nil {
		panic(errors.New("error parsing path"))
	}
	// url == "https://sanboxapi.rapid.net"
	return c.url.ResolveReference(endpoint).String()
}
