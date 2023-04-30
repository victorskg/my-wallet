package wallet

import "github.com/victorskg/my-wallet/internal/domain/stock"

type Dividend struct {
	info               stock.Dividend
	averagePriceOnDate float64
	quotasOnDate       uint32
	yeldOnCost         float64
	totalValue         float64
}
