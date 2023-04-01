package wallet

import (
	"github.com/victorskg/my-wallet/internal/domain/wallet"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type CreateWallet struct {
	walletGateway gateways.WalletGateway
}

func NewCreateWalletUseCase(walletGateway gateways.WalletGateway) CreateWallet {
	return CreateWallet{
		walletGateway: walletGateway,
	}
}

func (u CreateWallet) Execute(description string) (*wallet.Wallet, error) {
	wallet, err := wallet.NewWallet(description)
	if err != nil {
		return nil, err
	}

	return u.walletGateway.CreateWallet(wallet)
}
