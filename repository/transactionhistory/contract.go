package transactionhistory

import (
	"Julo/walletapp/entity"
	"context"
)

type Repository interface {
	InsertNewTransaction(ctx context.Context, trx entity.TransactionHistory) error
	GetAllTransactionByWalletID(ctx context.Context, walletID string) ([]entity.TransactionHistory, error)
}
