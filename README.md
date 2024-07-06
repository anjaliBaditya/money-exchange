# Money Exchange
=====================

A simple command-line Money Exchanger written in Go.

## Usage

To use the converter, run the program with three arguments:

* `amount`: the amount of money to convert
* `from_currency`: the currency to convert from (e.g. "EUR", "USD", etc.)
* `to_currency`: the currency to convert to (e.g. "GBP", "JPY", etc.)

Example:

``` bash
go run main.go 100 USD EUR
```

This would convert 100 US dollars to euros.

## API Endpoint

The converter uses the [ExchangeRate-API](https://exchangerate-api.com/) to fetch the latest exchange rates. You can replace the `apiURL` constant in the code with your own API endpoint if needed.

## Error Handling

The program will print an error message and exit if:

* The input amount is invalid
* The input currencies are invalid
* The API request fails
* The API response cannot be decoded

## Building and Running

To build and run the program, simply run `go build main.go` and then `./main` with the required arguments.

Note: This program assumes that the API endpoint returns exchange rates in the format expected by the `ExchangeRates` struct. If the API endpoint changes its response format, the program may need to be updated accordingly.
