package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const apiURL = "https://api.exchangerate-api.com/v4/latest/USD" // Replace with your chosen API endpoint

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func fetchRates() (ExchangeRates, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return ExchangeRates{}, fmt.Errorf("failed to fetch exchange rates: %w", err)
	}
	defer resp.Body.Close()

	var rates ExchangeRates
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return ExchangeRates{}, fmt.Errorf("failed to decode exchange rates: %w", err)
	}

	return rates, nil
}

func convert(amount float64, from string, to string, rates ExchangeRates) (float64, error) {
	fromRate, ok := rates.Rates[from]
	if !ok {
		return 0, fmt.Errorf("invalid currency code: %s", from)
	}

	toRate, ok := rates.Rates[to]
	if !ok {
		return 0, fmt.Errorf("invalid currency code: %s", to)
	}

	return amount / fromRate * toRate, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <amount> <from_currency> <to_currency>")
		return
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Invalid amount: %s\n", os.Args[1])
		return
	}

	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]

	rates, err := fetchRates()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	result, err := convert(amount, fromCurrency, toCurrency, rates)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}
