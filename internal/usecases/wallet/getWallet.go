package wallet

import (
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type GetWallet struct {
	walletGateway gateways.WalletGateway
}

func NewGetWalletUseCase(walletGateway gateways.WalletGateway) GetWallet {
	return GetWallet{walletGateway: walletGateway}
}

func (u GetWallet) Execute(id uuid.UUID) (*wallet.Wallet, error) {
	return u.walletGateway.GetWallet(id)
}
