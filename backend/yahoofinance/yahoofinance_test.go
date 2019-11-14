package yahoofinance

import (
	"github.com/NicksonT/StockSim/backend/stock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetQuoteOf(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()
	expected := stock.Stock{Symbol: "MSFT", Price: "146.1100"}
	actual := stock.Stock{}
	mock := YahooFinance{URL: ts.URL}
	mock.GetQuoteOf("MSFT", &actual)
	assert.Equal(t, expected, actual)
}
func newTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/MSFT", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html>
  <div id="quote-header-info">
    <div>
    </div>
    <div>
    </div>
    <div>
      <div>
        <span>146.1100</span>
      </div>
    </div>
  </div>
</html>`))
	})
	return httptest.NewServer(mux)
}
