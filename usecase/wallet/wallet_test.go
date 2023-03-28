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

func TestModule_CheckWalletBalance(t *testing.T) {
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

	wlt, err := m.CheckWalletBalance(ctx, "new-token")
	if err != nil {
		panic(err)
	}
	fmt.Println(wlt)
}

func TestModule_CheckWalletTransactionHistory(t *testing.T) {
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

	trx, err := m.CheckWalletTransactionHistory(ctx, "new-token")
	if err != nil {
		panic(err)
	}
	fmt.Println(trx)
}

func TestModule_updateWalletBalance(t *testing.T) {
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

	wlt, err := m.updateWalletBalance(ctx, UpdateWalletBalanceReq{
		Token:       "new-token",
		Amount:      -50000,
		ReferenceID: uuid.NewString(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(wlt)
}
