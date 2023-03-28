package gateway

import "github.com/victorskg/my-wallet/internal/domain/stock"

type StockGateway interface {
	SaveStocks(stocks []stock.Stock) ([]stock.Stock, error)
}
