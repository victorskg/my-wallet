package database

import (
	"github.com/victorskg/my-wallet/internal/domain/stock"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type StockDatabaseGateway struct {
}

func NewStockDatabaseGateway() gateways.StockGateway {
	return StockDatabaseGateway{}
}

func (g StockDatabaseGateway) SaveStocks(stocks []stock.Stock) ([]stock.Stock, error) {
	return stocks, nil
}

func (g StockDatabaseGateway) FindByTicker(ticker string) (*stock.Stock, error) {
	return nil, nil
}
