package wallet

import (
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
	"github.com/victorskg/my-wallet/internal/usecases/wallet/gateway"
)

type GetWallet struct {
	walletGateway gateway.WalletGateway
}

func (u GetWallet) Execute(id uuid.UUID) (*wallet.Wallet, error) {
	return u.walletGateway.GetWallet(id)
}
