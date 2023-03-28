package wallet

import (
	"Julo/walletapp/entity"
	"context"
)

type Usecase interface {
	InitAccountWallet(ctx context.Context, customerID string) (string, error)
	ChangeWalletStatus(ctx context.Context, token string, isEnable bool) (entity.Wallet, error)
	CheckWalletBalance(ctx context.Context, token string) (entity.Wallet, error)
	CheckWalletTransactionHistory(ctx context.Context, token string) ([]entity.TransactionHistory, error)
	DepositMoneyToWallet(ctx context.Context, req UpdateWalletBalanceReq) (entity.Wallet, error)
	WithdrawMoneyFromWallet(ctx context.Context, req UpdateWalletBalanceReq) (entity.Wallet, error)
}
