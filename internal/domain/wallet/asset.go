package wallet

import "time"

type Investment struct {
	Quotas   int16
	BuyDate  time.Time
	BuyPrice float64
}

type Asset struct {
	ticker       string
	averagePrice float64
	investments  []Investment
	dividends    []Dividend
}

func NewAsset(ticker string) *Asset {
	return &Asset{
		ticker: ticker,
	}
}

func (a *Asset) Invest(investment Investment) {
	a.investments = append(a.investments, investment)
	a.calculateAveragePrice()
}

func (a *Asset) TotalQuotas() int16 {
	var totalQuotas int16 = 0
	for _, investment := range a.investments {
		totalQuotas += investment.Quotas
	}
	return totalQuotas
}

func (a *Asset) TotalPrice() float64 {
	return float64(a.TotalQuotas()) * a.averagePrice
}

func (a *Asset) calculateAveragePrice() {
	var totalPaid float64 = 0
	for _, investment := range a.investments {
		totalPaid += investment.BuyPrice
	}

	a.averagePrice = totalPaid / float64(a.TotalQuotas())
}
