package wallet

import (
	"time"

	"github.com/victorskg/my-wallet/internal/domain/stock"
)

type Investment struct {
	Quotas   uint32
	BuyDate  time.Time
	BuyPrice float64
}

type AveragePriceHistory struct {
	averagePrice float64
	date         time.Time
}

type Asset struct {
	ticker              string
	averagePrice        float64
	averagePriceHistory []AveragePriceHistory
	investments         []Investment
	dividends           []Dividend
}

func NewAsset(ticker string) *Asset {
	return &Asset{
		ticker: ticker,
	}
}

func (a *Asset) Ticker() string {
	return a.ticker
}

func (a *Asset) Invest(investment Investment) {
	a.investments = append(a.investments, investment)
	a.calculateAveragePrice(investment.BuyDate)
}

func (a *Asset) TotalQuotas() uint32 {
	var totalQuotas uint32 = 0
	for _, investment := range a.investments {
		totalQuotas += investment.Quotas
	}
	return totalQuotas
}

func (a *Asset) TotalPrice() float64 {
	return float64(a.TotalQuotas()) * a.averagePrice
}

func (a *Asset) AddDividends(dividends ...stock.Dividend) {
	uniqueDividends := make(map[stock.Dividend]bool)
	for _, value := range a.dividends {
		if !uniqueDividends[value.info] {
			uniqueDividends[value.info] = true
		}
	}

	for _, stockDividend := range dividends {
		if !uniqueDividends[stockDividend] {
			assetDividend := a.createAssetDividend(stockDividend)
			a.dividends = append(a.dividends, assetDividend)
		}
	}
}

func (a *Asset) calculateAveragePrice(date time.Time) {
	var totalPaid float64 = 0
	for _, investment := range a.investments {
		totalPaid += investment.BuyPrice
	}

	a.averagePrice = totalPaid / float64(a.TotalQuotas())
	a.saveAveragePriceOnHistory(date)
}

func (a *Asset) saveAveragePriceOnHistory(date time.Time) {
	averagePriceHistorySize := len(a.averagePriceHistory)
	if averagePriceHistorySize > 0 {
		lastAveragePrice := &a.averagePriceHistory[averagePriceHistorySize-1]
		if lastAveragePrice.date == date {
			lastAveragePrice.averagePrice = a.averagePrice
			return
		}
	}

	a.averagePriceHistory = append(a.averagePriceHistory, AveragePriceHistory{averagePrice: a.averagePrice, date: date})
}

func (a *Asset) createAssetDividend(dividend stock.Dividend) Dividend {
	baseDate := dividend.BaseDate()
	averagePriceOnDate := a.averagePriceOnDate(baseDate)
	quotasOnDate := a.quotasOnDate(baseDate)
	yeldOnCost := (dividend.Value() / averagePriceOnDate) * 100
	totalValue := float64(quotasOnDate) * dividend.Value()

	return Dividend{
		info:               dividend,
		averagePriceOnDate: averagePriceOnDate,
		quotasOnDate:       a.quotasOnDate(baseDate),
		yeldOnCost:         yeldOnCost,
		totalValue:         totalValue,
	}
}

func (a *Asset) averagePriceOnDate(baseDate time.Time) float64 {
	var smallestDiff time.Duration
	var averagePriceBeforeDividend *AveragePriceHistory
	for _, aph := range a.averagePriceHistory {
		if averagePriceBeforeDividend == nil {
			averagePriceBeforeDividend = &aph
		} else {
			diff := aph.date.Sub(baseDate)
			if aph.date.Before(baseDate) &&
				(smallestDiff == 0 || diff < smallestDiff) {
				smallestDiff = diff
				averagePriceBeforeDividend = &aph
			}
		}
	}

	return averagePriceBeforeDividend.averagePrice
}

func (a *Asset) quotasOnDate(baseDate time.Time) uint32 {
	var quotasOnDate uint32 = 0
	for _, investment := range a.investments {
		if investment.BuyDate.Before(baseDate) {
			quotasOnDate += investment.Quotas
		}
	}

	return quotasOnDate
}
