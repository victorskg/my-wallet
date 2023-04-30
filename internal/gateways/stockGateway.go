package gateways

import "github.com/victorskg/my-wallet/internal/domain/stock"

type StockGateway interface {
	SaveStocks(stocks []stock.Stock) ([]stock.Stock, error)
	FindByTicker(ticker string) (*stock.Stock, error)
	UpdateStock(stockDomain *stock.Stock) (*stock.Stock, error)
}
