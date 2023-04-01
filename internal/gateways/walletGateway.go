package gateways

import (
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
)

type WalletGateway interface {
	CreateWallet(wallet *wallet.Wallet) (*wallet.Wallet, error)
	GetWallet(id uuid.UUID) (*wallet.Wallet, error)
	UpdateWallet(wallet *wallet.Wallet) (*wallet.Wallet, error)
}
