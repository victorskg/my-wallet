package stock

import "time"

type SType string

const (
	Common     SType = "common"
	RealEstate SType = "real_estate"
)

type Price struct {
	date  time.Time
	price float64
}

type Stock struct {
	ticker           string
	name             string
	sType            SType
	category         string
	subCategory      string
	administrator    string
	bookValue        float32
	patrimony        float64
	pvp              float32
	dividends        []Dividend
	historicalPrices []Price
}

func NewStock(ticker string, sType SType, category string, subCategory string,
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

func (s *Stock) AddDividends(d ...Dividend) {
	uniqueDividends := make(map[Dividend]bool)
	for _, value := range s.dividends {
		if !uniqueDividends[value] {
			uniqueDividends[value] = true
		}
	}

	for _, value := range d {
		if !uniqueDividends[value] {
			s.dividends = append(s.dividends, value)
		}
	}
}
