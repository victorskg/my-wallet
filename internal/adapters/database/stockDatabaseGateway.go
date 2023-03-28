package database

import (
	"github.com/victorskg/my-wallet/internal/domain/stock"
	"github.com/victorskg/my-wallet/internal/usecases/stock/gateway"
)

type StockDatabaseGateway struct {
}

func NewStockDatabaseGateway() gateway.StockGateway {
	return StockDatabaseGateway{}
}

func (g StockDatabaseGateway) SaveStocks(stocks []stock.Stock) ([]stock.Stock, error) {
	return stocks, nil
}
