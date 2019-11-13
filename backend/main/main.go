package main

import (
	"fmt"
	"net/http"
)

func main() {
	rat := backend.AlphaVantage{}
	rat.APIKey = "87HDEBBD5YXA73ZW"
	stock := backend.Stock{}
	c := http.Client{}
	rat.GetQuoteOf("TSLA", &stock, c)
	fmt.Print(stock)
}
