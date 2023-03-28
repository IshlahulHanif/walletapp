package api

import "Julo/walletapp/usecase/wallet"

type Module struct {
	usecase usecase
}

type usecase struct {
	wallet wallet.Usecase
}

type TransactionsHistory struct {
	Transaction []Transaction `json:"transaction"`
}

type Transaction struct {
	ID           string  `json:"id"`
	Status       string  `json:"status"`
	TransactedAt string  `json:"transacted_at"`
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	RefID        string  `json:"Reference_id"`
}
