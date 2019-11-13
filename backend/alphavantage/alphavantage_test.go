package alphavantage

import (
	"github.com/karupanerura/go-mock-http-response"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuoteOf(t *testing.T) {
	mockResponse := []byte(`{
    "Global Quote": {
        "01. symbol": "MSFT",
        "02. open": "145.3400",githu
        "03. high": "146.4200",
        "04. low": "144.7300",
        "05. price": "146.1100",
        "06. volume": "13042200",
        "07. latest trading day": "2019-11-11",
        "08. previous close": "145.9600",
        "09. change": "0.1500",
        "10. change percent": "0.1028%"
    }
}`)

	expected := Stock{Symbol: "MSFT", Price: "146.1100"}
	actual := Stock{}
	mock := AlphaVantage{}
	c := mockhttp.NewResponseMock(200, nil, mockResponse).MakeClient()
	mock.GetQuoteOf("MSFT", &actual, *c)
	assert.Equal(t, expected, actual)
}

func TestRemoveGlobalQuote(t *testing.T) {
	mockResponse := []byte(`{
    "Global Quote": {
        "01. symbol": "MSFT",
        "02. open": "145.3400",
        "03. high": "146.4200",
        "04. low": "144.7300",
        "05. price": "146.1100",
        "06. volume": "13042200",
        "07. latest trading day": "2019-11-11",
        "08. previous close": "145.9600",
        "09. change": "0.1500",
        "10. change percent": "0.1028%"
    }
}`)
	expected := []byte(`{
            "01. symbol": "MSFT",
        "02. open": "145.3400",
        "03. high": "146.4200",
        "04. low": "144.7300",
        "05. price": "146.1100",
        "06. volume": "13042200",
        "07. latest trading day": "2019-11-11",
        "08. previous close": "145.9600",
        "09. change": "0.1500",
        "10. change percent": "0.1028%"
    }`)

	assert.Equal(t, string(expected), string(RemoveGlobalQuote(mockResponse)))
}
