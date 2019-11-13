package alphavantage

import (
	"bytes"
	"encoding/json"
	"github.com/NicksonT/StockSim/backend/stock"
	"io/ioutil"
	"log"
	"net/http"
)

type AlphaVantage struct {
	APIKey string
}

func (m *AlphaVantage) GetQuoteOf(s string, st *stock.Stock, c http.Client) {
	req, err := http.NewRequest("GET", "https://www.alphavantage.co/query?function=GLOBAL_QUOTE", nil)
	if err != nil {
		log.Print(err)
	}

	q := req.URL.Query()
	q.Add("apikey", m.APIKey)
	q.Add("symbol", s)
	req.URL.RawQuery = q.Encode()
	resp, err := c.Do(req)
	body := resp.Body
	b, _ := ioutil.ReadAll(body)
	err = json.Unmarshal(RemoveGlobalQuote(b), st)
	if err != nil {
		log.Print(err)
	}
}

func RemoveGlobalQuote(b []byte) []byte {
	b = bytes.Replace(b, []byte("\"Global Quote\": {\n"), []byte(""), 1)
	b = bytes.Replace(b, []byte("}\n"), []byte(""), 1)
	return b
}
