package backend

type StockExchange struct {
	stocks map[string]Stock
}

func (s *StockExchange) ViewInfo(st string) Stock {
	return s.stocks[st]
}
