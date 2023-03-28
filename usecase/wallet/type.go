package wallet

import (
	"Julo/walletapp/repository/transactionhistory"
	"Julo/walletapp/repository/userprovider"
	"Julo/walletapp/repository/wallet"
)

type Module struct {
	repo repo
}

type repo struct {
	transaction transactionhistory.Repository
	user        userprovider.Repository
	wallet      wallet.Repository
}

type UpdateWalletBalanceReq struct {
	Token       string
	Amount      float64
	ReferenceID string
}
