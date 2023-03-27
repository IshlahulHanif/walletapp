package wallet

import (
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestModule_InitAccountWallet(t *testing.T) {
	ctx := context.Background()

	m, err := GetInstance(utils.Config{
		DatabaseConfig: database.ConfigDatabase{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			Dbname:   "wallet_db",
		},
	})
	if err != nil {
		panic(err)
	}

	custID := uuid.NewString()

	token, err := m.InitAccountWallet(ctx, custID)
	if err != nil {
		return
	}
	fmt.Println(token)
}

func TestModule_ChangeWalletStatus(t *testing.T) {
	ctx := context.Background()

	m, err := GetInstance(utils.Config{
		DatabaseConfig: database.ConfigDatabase{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			Dbname:   "wallet_db",
		},
	})
	if err != nil {
		panic(err)
	}

	wlt, err := m.ChangeWalletStatus(ctx, "new-token", false)
	if err != nil {
		panic(err)
	}
	fmt.Println(wlt)
}
