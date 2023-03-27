package wallet

import (
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"context"
	"fmt"
	"testing"
)

func TestModule_CreateNewWalletCreateWalletTest(t *testing.T) {
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

	err = m.CreateNewWallet(ctx, "dc4fc339-787f-4d6b-b197-686065ca8959")
	if err != nil {
		panic(err)
	}
}

func TestModule_UpdateWalletAmountTest(t *testing.T) {
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

	err = m.UpdateWalletAmount(ctx, "dc4fc339-787f-4d6b-b197-686065ca8959", 10000)
	if err != nil {
		panic(err)
	}
}

func TestModule_IncrWalletAmountTest(t *testing.T) {
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

	err = m.IncrWalletAmount(ctx, "dc4fc339-787f-4d6b-b197-686065ca8959", 10000)
	if err != nil {
		panic(err)
	}
}

func TestModule_GetWalletAmountTest(t *testing.T) {
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

	amount, err := m.GetWalletAmountByCustomerID(ctx, "dc4fc339-787f-4d6b-b197-686065ca8959")
	if err != nil {
		panic(err)
	}
	fmt.Println(amount)
}

func TestModule_UpdateWalletStatusByCustomerIDTest(t *testing.T) {
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

	err = m.UpdateWalletStatusByCustomerID(ctx, "dc4fc339-787f-4d6b-b197-686065ca8959", true)
	if err != nil {
		panic(err)
	}
}
