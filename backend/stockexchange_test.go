package backend

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewInfoAboutStock(t *testing.T) {
	stocks := make(map[string]Stock)
	stocks["AAPL"] = Stock{"Apple", 234.30}
	nyse := StockExchange{stocks: stocks}
	assert.Equal(t, "Apple", nyse.ViewInfo("AAPL").name)
}
