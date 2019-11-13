package main

import (
	"fmt"
	"github.com/NicksonT/StockSim/backend/alphavantage"
	"github.com/NicksonT/StockSim/backend/stock"
	"net/http"
)

func main() {
	rat := alphavantage.AlphaVantage{}
	rat.APIKey = "87HDEBBD5YXA73ZW"
	stock := stock.Stock{}
	c := http.Client{}
	rat.GetQuoteOf("TSLA", &stock, c)
	fmt.Print(stock)
}
