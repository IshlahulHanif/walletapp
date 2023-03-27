package wallet

import (
	"Julo/walletapp/entity"
	"context"
)

type Repository interface {
	CreateNewWallet(ctx context.Context, customerID string) error
	UpdateWalletAmount(ctx context.Context, customerID string, amount float64) error
	IncrWalletAmount(ctx context.Context, customerID string, amount float64) error
	GetWalletAmountByCustomerID(ctx context.Context, customerID string) (entity.Wallet, error)
	UpdateWalletStatusByCustomerID(ctx context.Context, customerID string, isEnabled bool) error
}

type MemoryMethod interface {
}
