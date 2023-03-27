package wallet

import "Julo/walletapp/pkg/database"

type Module struct {
	memory   MemoryMethod
	database database.MethodDatabase
}

type MemoryResource struct { // TODO: delete if we wont implement mem cache
	UserWallet map[string]Wallet
}

type Wallet struct { // TODO: delete if we wont implement mem cache
	IsEnabled bool    `json:"isEnabled"`
	Amount    float64 `json:"amount"`
}
