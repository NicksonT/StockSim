package stockexchange

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewInfoAboutStock(t *testing.T) {
	stocks := make(map[string]Stock)
	stocks["AAPL"] = Stock{"Apple", "AAPL", "234.30"}
	nyse := StockExchange{stocks: stocks}
	AAPLStock, _ := nyse.ViewInfo("AAPL")
	assert.Equal(t, "Apple", AAPLStock.name)
}

func TestErrorReturnedForStockNotFound(t *testing.T) {
	stocks := make(map[string]Stock)
	stocks["AAPL"] = Stock{"Apple", "AAPL", "234.30"}
	nyse := StockExchange{stocks: stocks}
	_, err := nyse.ViewInfo("ZE")
	assert.Error(t, errors.New("This stock can not be found in the exchange"), err)
}
