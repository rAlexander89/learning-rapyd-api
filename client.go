package main

import (
	"errors"
	"net/http"
	"net/url"
)

type client struct {
	Signer
	*http.Client
	url *url.URL
}

type Client interface {
	Sign(path string) ([]byte, error)
	GetCountries() (*CountriesResponse, error)
	Resolve(path string) string
}

func NewClient(signer Signer, cli *http.Client, url *url.URL) Client {
	return &client{
		Signer: signer,
		Client: cli,
		url:    url,
	}
}

func (c *client) Resolve(path string) string {
	endpoint, err := url.Parse(path)
	if err != nil {
		panic(errors.New("error parsing path"))
	}
	// url == "https://sanboxapi.rapid.net"
	return c.url.ResolveReference(endpoint).String()
}
