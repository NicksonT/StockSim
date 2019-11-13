package stockexchange

import (
	"errors"
)

type StockExchange struct {
	stocks map[string]Stock
}

func (s *StockExchange) ViewInfo(st string) (*Stock, error) {

	stock, found := s.stocks[st]
	if found != true {
		return nil, errors.New("This stock can not be found in the exchange")
	}
	return &stock, nil
}
