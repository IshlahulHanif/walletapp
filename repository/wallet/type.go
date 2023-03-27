package wallet

import "Julo/walletapp/pkg/database"

type Module struct {
	memory   MemoryMethod
	database database.MethodDatabase
}

type MemoryResource struct {
	UserWallet map[string]Wallet
}

type Wallet struct {
	IsEnabled bool    `json:"isEnabled"`
	Amount    float64 `json:"amount"`
}
