package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

const (
	getCountries = "v1/data/countries"
)

type Data interface{}

type CountriesResponse struct {
	Data Data `json:"data"`
}

func (c *client) GetCountries() (*CountriesResponse, error) {
	res, err := c.Sign(getCountries)
	if err != nil {
		return nil, errors.New("error getting countries")
	}

	var body CountriesResponse

	err = json.Unmarshal(res, &body)

	if err != nil {
		log.Fatal(err)
		fmt.Print("err with json shit")
		return nil, errors.New("error unmarshalleing response")
	}

	return &body, nil
}
