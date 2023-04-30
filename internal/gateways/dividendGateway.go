package gateways

import (
	"github.com/victorskg/my-wallet/internal/domain/stock"
)

type DividendGateway interface {
	ImportDividends(ticker string) ([]stock.Dividend, error)
}
