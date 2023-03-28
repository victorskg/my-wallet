package stock

import "time"

type StockType string

const (
	Common    StockType = "common"
	RealState StockType = "real_state"
)

type Price struct {
	date  time.Time
	price float64
}

type Stock struct {
	ticker           string
	name             string
	sType            StockType
	category         string
	subCategory      string
	administrator    string
	bookValue        float32
	patrimony        float64
	pvp              float32
	dividends        []Dividend
	historicalPrices []Price
}

func NewStock(ticker string, sType StockType, category string, subCategory string,
	administrator string, bookValue float32, patrimony float64, pvp float32) *Stock {
	return &Stock{
		ticker:        ticker,
		sType:         sType,
		category:      category,
		subCategory:   subCategory,
		administrator: administrator,
		bookValue:     bookValue,
		patrimony:     patrimony,
		pvp:           pvp,
	}
}
