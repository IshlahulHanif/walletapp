package entity

import (
	"time"
)

type TransactionHistory struct {
	ID           string    `json:"id" db:"id"`
	Status       string    `json:"status" db:"status"`
	WalletID     int64     `json:"wallet_id" db:"wallet_id"`
	TransactedAt time.Time `json:"transacted_at" db:"transacted_at"`
	Type         string    `json:"type" db:"type"`
	Amount       float64   `json:"amount" db:"amount"`
	ReferenceID  string    `json:"reference_id" db:"reference_id"`
}
