package wallet

import "github.com/victorskg/my-wallet/internal/domain/stock"

type Dividend struct {
	info               stock.Dividend
	averagePriceOnDate float64
	quotasOnDate       float64
	yeldOnCost         float32
	totalValue         float64
}
