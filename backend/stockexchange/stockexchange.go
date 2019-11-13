package stockexchange

import (
	"errors"
	"github.com/NicksonT/StockSim/backend/stock"
)

type StockExchange struct {
	stocks map[string]stock.Stock
}

func (s *StockExchange) ViewInfo(st string) (*stock.Stock, error) {

	stock, found := s.stocks[st]
	if found != true {
		return nil, errors.New("This stock can not be found in the exchange")
	}
	return &stock, nil
}
