package main

import (
	"fmt"
	"log"
)

func main() {
	ratesData, err := FetchRates()
	if err != nil {
		log.Fatalf("Failed to fetch rates: %v", err)
	}

	amount := 100.0
	from := "USD"
	to := "EUR"

	result, err := ConvertCurrency(amount, from, to, ratesData.Rates)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}
