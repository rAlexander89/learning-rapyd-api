package rapyd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/rAlexander89/rapyd-banking/countries"
)

func (c *RapydClient) GetCountries() (*countries.GetCountriesResponse, error) {
	res, err := c.Sign(countries.Index)
	if err != nil {
		return nil, errors.New("error getting countries")
	}

	var body countries.GetCountriesResponse

	err = json.Unmarshal(res, &body)

	if err != nil {
		log.Fatal(err)
		fmt.Print("err with json shit")
		return nil, errors.New("error unmarshalleing response")
	}

	return &body, nil
}
