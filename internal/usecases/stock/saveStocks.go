package stock

import (
	"github.com/victorskg/my-wallet/internal/domain/stock"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type SaveStock struct {
	stockGateway gateways.StockGateway
}

func NewSaveStockUseCase(stockGateway gateways.StockGateway) SaveStock {
	return SaveStock{
		stockGateway: stockGateway,
	}
}

func (u SaveStock) Execute(stocks ...stock.Stock) ([]stock.Stock, error) {
	return u.stockGateway.SaveStocks(stocks)
}
