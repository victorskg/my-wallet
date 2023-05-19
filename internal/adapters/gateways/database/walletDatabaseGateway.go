package database

import (
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/victorskg/my-wallet/internal/adapters/gateways/database/entity"
	"github.com/victorskg/my-wallet/internal/adapters/gateways/database/mapper"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
	"github.com/victorskg/my-wallet/internal/gateways"
	"github.com/victorskg/my-wallet/pkg/database"
	pkgError "github.com/victorskg/my-wallet/pkg/error"
)

type WalletDatabaseGateway struct {
	repository *database.Repository[entity.WalletEntity]
	mapper     mapper.Mapper[entity.WalletEntity, wallet.Wallet]
}

func NewWalletDatabaseGateway(repository *database.Repository[entity.WalletEntity]) gateways.WalletGateway {
	return WalletDatabaseGateway{repository: repository, mapper: mapper.WalletMapper()}
}

func (g WalletDatabaseGateway) CreateWallet(wallet *wallet.Wallet) (*wallet.Wallet, error) {
	return wallet, nil
}

func (g WalletDatabaseGateway) GetWallet(id uuid.UUID) (*wallet.Wallet, error) {
	if err := g.repository.Connect(); err != nil {
		return nil, err
	}

	walletEntity, err := g.repository.SelectOne(fmt.Sprintf("id = '%s'", id.String()))
	if err != nil {
		return nil, err
	}

	if walletEntity == nil && err == nil {
		return nil, pkgError.NewNotFoundError(id)
	}

	walletDomain := g.mapper.FromEntityToDomain(*walletEntity)
	return walletDomain, nil
}

func (g WalletDatabaseGateway) UpdateWallet(wallet *wallet.Wallet) (*wallet.Wallet, error) {
	return nil, nil
}
