package entity

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type WalletEntity struct {
	id          uuid.UUID
	description string
}

func (w WalletEntity) Id() uuid.UUID {
	return w.id
}

func (w WalletEntity) Description() string {
	return w.description
}

func NewWalletEntity(id uuid.UUID, description string) (*WalletEntity, error) {
	return validate(&WalletEntity{id: id, description: description})
}

func validate(wallet *WalletEntity) (*WalletEntity, error) {
	if wallet.description == "" {
		return nil, errors.New("A descrição da carteira é obrigatória.")
	}

	return wallet, nil
}

func (w WalletEntity) FromRow(rows *sql.Rows) (any, error) {
	var id uuid.UUID
	var description string

	if err := rows.Scan(&id, &description); err != nil {
		return nil, err
	}

	return NewWalletEntity(id, description)
}
