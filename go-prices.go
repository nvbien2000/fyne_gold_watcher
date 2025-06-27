package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var currency = "USD"

// Gold -
type Gold struct {
	Items  []PriceItem `json:"items"`
	Client *http.Client
}

// PriceItem -
type PriceItem struct {
	Currency      string  `json:"curr"`
	GoldPrice     float64 `json:"xauPrice"`
	SilverPrice   float64 `json:"xagPrice"`
	GoldChange    float64 `json:"chgXau"`
	SilverChange  float64 `json:"chgXag"`
	GoldPercent   float64 `json:"pcXau"`
	SilverPercent float64 `json:"pcXag"`
	GoldClose     float64 `json:"xauClose"`
	SilverClose   float64 `json:"xagClose"`
}

// GetPrices -
func (g *Gold) GetPrices() (*PriceItem, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client
	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print("error creating new request to goldprice.org", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Print("error creating new request to goldprice.org", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading json", err)
		return nil, err
	}

	gold := Gold{}
	// var prev, curr, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("unmarshal error", err)
		return nil, err
	}

	currentInfo := gold.Items[0]
	return &currentInfo, nil
}
