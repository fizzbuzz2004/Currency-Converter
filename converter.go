package main

import "fmt"

func ConvertCurrency(amount float64, from, to string, rates map[string]float64) (float64, error) {
	fromRate, ok := rates[from]
	if !ok {
		return 0, fmt.Errorf("invalid source currency: %s", from)
	}

	toRate, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("invalid target currency: %s", to)
	}

	amountInUSD := amount / fromRate
	return amountInUSD * toRate, nil
}
