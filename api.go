package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ExchangeRates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func FetchRates() (*ExchangeRates, error) {
	appID := os.Getenv("OER_APP_ID")
	if appID == "" {
		return nil, fmt.Errorf("OER_APP_ID environment variable not set")
	}

	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", appID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status code %d", resp.StatusCode)
	}

	var rates ExchangeRates
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, err
	}

	return &rates, nil
}
