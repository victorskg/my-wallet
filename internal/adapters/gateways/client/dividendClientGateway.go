package client

import (
	"github.com/victorskg/my-wallet/internal/domain/stock"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type DividendClientGateway struct {
}

func NewDividendClientGateway() gateways.DividendGateway {
	return DividendClientGateway{}
}

func (d DividendClientGateway) ImportDividends(ticker string) ([]stock.Dividend, error) {
	return nil, nil
}
