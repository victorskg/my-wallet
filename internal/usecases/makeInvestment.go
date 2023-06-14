package usecases

import (
	"time"

	"github.com/victorskg/my-wallet/internal/gateways"

	domain "github.com/victorskg/my-wallet/internal/domain/wallet"

	"github.com/google/uuid"
)

type MakeInvestmentInput struct {
	WalletID uuid.UUID
	Assets   []InvestmentAsset
}

type InvestmentAsset struct {
	Ticker  string
	Price   float64
	Quotas  uint32
	BuyDate time.Time
}

type MakeInvestment struct {
	walletGateway gateways.WalletGateway
	stockGateway  gateways.StockGateway
}

func NewMakeInvestmentUseCase(walletGateway gateways.WalletGateway, stockGateway gateways.StockGateway) MakeInvestment {
	return MakeInvestment{
		walletGateway: walletGateway,
		stockGateway:  stockGateway,
	}
}

func (u MakeInvestment) Execute(input MakeInvestmentInput) error {
	wallet, err := u.walletGateway.GetWallet(input.WalletID)
	if err != nil {
		return err
	}

	for _, asset := range input.Assets {
		_, tickerErr := u.stockGateway.FindByTicker(asset.Ticker)
		if tickerErr != nil {
			return tickerErr
		}

		//TODO: Create constructor and validate values
		investment := domain.Investment{
			Quotas:   asset.Quotas,
			BuyDate:  asset.BuyDate,
			BuyPrice: asset.Price,
		}
		wallet.MakeInvestment(asset.Ticker, investment)
	}

	_, err = u.walletGateway.UpdateWallet(wallet)
	if err != nil {
		return err
	}

	return nil
}
