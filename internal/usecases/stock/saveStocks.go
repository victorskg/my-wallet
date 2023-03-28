package stock

import (
	"github.com/victorskg/my-wallet/internal/domain/stock"
	"github.com/victorskg/my-wallet/internal/usecases/stock/gateway"
)

type SaveStock struct {
	stockGateway gateway.StockGateway
}

func NewSaveStockUseCase(stockGateway gateway.StockGateway) SaveStock {
	return SaveStock{
		stockGateway: stockGateway,
	}
}

func (u SaveStock) Execute(stocks ...stock.Stock) ([]stock.Stock, error) {
	return u.stockGateway.SaveStocks(stocks)
}
