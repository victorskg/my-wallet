package wallet

import (
	"time"

	domain "github.com/victorskg/my-wallet/internal/domain/wallet"

	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/usecases/wallet/gateway"
)

type Input struct {
	walletId uuid.UUID
	assets   []struct {
		ticker  string
		price   float64
		quotas  int16
		buyDate time.Time
	}
}

type MakeInvestment struct {
	walletGateway gateway.WalletGateway
}

func (u MakeInvestment) Execute(input Input) error {
	wallet, err := u.walletGateway.GetWallet(input.walletId)
	if err != nil {
		return err
	}

	for _, asset := range input.assets {
		//TODO: Create constructor and validate values
		investment := domain.Investment{
			Quotas:   asset.quotas,
			BuyDate:  asset.buyDate,
			BuyPrice: asset.price,
		}
		wallet.MakeInvestment(asset.ticker, investment)
	}

	_, err = u.walletGateway.UpdateWallet(wallet)
	if err != nil {
		return err
	}

	return nil
}
