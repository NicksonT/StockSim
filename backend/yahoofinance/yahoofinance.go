package yahoofinance

import (
	"github.com/NicksonT/StockSim/backend/stock"
	"github.com/gocolly/colly"
	"log"
)

type YahooFinance struct {
	URL string
}

func (y *YahooFinance) GetQuoteOf(sy string, st *stock.Stock) {
	c := colly.NewCollector()
	syURL := y.URL + "/" + sy
	c.OnXML("//*[@id=\"quote-header-info\"]/div[3]/div[1]/span", func(e *colly.XMLElement) {
		st.Price = e.Text
		st.Symbol = sy
	})
	err := c.Visit(syURL)
	if err != nil {
		log.Print(err)
	}
}
