package database

import (
	"github.com/google/uuid"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
	"github.com/victorskg/my-wallet/internal/gateways"
)

type WalletDatabaseGateway struct {
}

func NewWalletDatabaseGateway() gateways.WalletGateway {
	return WalletDatabaseGateway{}
}

func (g WalletDatabaseGateway) CreateWallet(wallet *wallet.Wallet) (*wallet.Wallet, error) {
	return wallet, nil
}

func (g WalletDatabaseGateway) GetWallet(id uuid.UUID) (*wallet.Wallet, error) {
	return nil, nil
}

func (g WalletDatabaseGateway) UpdateWallet(wallet *wallet.Wallet) (*wallet.Wallet, error) {
	return nil, nil
}
