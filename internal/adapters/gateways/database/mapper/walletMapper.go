package mapper

import (
	"github.com/victorskg/my-wallet/internal/adapters/gateways/database/entity"
	"github.com/victorskg/my-wallet/internal/domain/wallet"
)

type walletMapper struct {
}

func WalletMapper() walletMapper {
	return walletMapper{}
}

func (w walletMapper) FromEntityToDomain(wEntity entity.WalletEntity) *wallet.Wallet {
	wDomain, _ := wallet.NewWalletFromDatabase(wEntity.Id(), wEntity.Description())
	return wDomain
}

func (w walletMapper) FromDomainToEntity(wDomain wallet.Wallet) *entity.WalletEntity {
	wEntity, _ := entity.NewWalletEntity(wDomain.Id(), wDomain.Description())
	return wEntity
}
