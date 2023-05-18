package wallet

import (
	"errors"
	"github.com/google/uuid"
)

type tickerPrice map[string]float64

type Wallet struct {
	id          uuid.UUID
	description string
	assets      []Asset
}

func NewWallet(description string) (*Wallet, error) {
	return validate(&Wallet{
		id:          uuid.New(),
		description: description,
	})
}

func NewWalletFromDatabase(id uuid.UUID, description string) (*Wallet, error) {
	return validate(&Wallet{
		id:          id,
		description: description,
	})
}

func validate(wallet *Wallet) (*Wallet, error) {
	if wallet.description == "" {
		return nil, errors.New("A descrição da carteira é obrigatória.")
	}

	return wallet, nil
}

func (w Wallet) Id() uuid.UUID {
	return w.id
}

func (w Wallet) Description() string {
	return w.description
}

func (w Wallet) Assets() []Asset {
	return w.assets
}

func (w Wallet) MakeInvestment(ticker string, investment Investment) {
	var asset *Asset
	for _, a := range w.assets {
		if a.ticker == ticker {
			asset = &a
			break
		}
	}

	if asset == nil {
		asset = NewAsset(ticker)
		w.assets = append(w.assets, *asset)
	}

	asset.Invest(investment)
}

func (w Wallet) TotalInvested() float64 {
	totalInvested := 0.0
	for _, asset := range w.assets {
		totalInvested += asset.TotalPrice()
	}
	return totalInvested
}

func (w Wallet) CurrentPatrimony(currentTickerPrice tickerPrice) float64 {
	patrimony := 0.0
	for _, asset := range w.assets {
		value := currentTickerPrice[asset.ticker]
		patrimony += value * float64(asset.TotalQuotas())
	}
	return patrimony
}

func (w Wallet) PatrimonyFloatationPercentage(currentTickerPrice tickerPrice) float64 {
	return ((w.TotalInvested() / w.CurrentPatrimony(currentTickerPrice)) - 1) * 100
}
