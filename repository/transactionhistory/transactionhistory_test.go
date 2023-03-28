package transactionhistory

import (
	"Julo/walletapp/entity"
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestModule_GetAllTransactionByWalletID(t *testing.T) {
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

	res, err := m.GetAllTransactionByWalletID(ctx, "42487142-1ee7-4326-94b9-c3d2386e493a")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestModule_InsertNewTransaction(t *testing.T) {
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

	trx := entity.TransactionHistory{
		Status:       "success",
		WalletID:     "42487142-1ee7-4326-94b9-c3d2386e493a",
		TransactedAt: time.Now(),
		Type:         "deposit",
		Amount:       4000,
		ReferenceID:  uuid.NewString(),
	}

	err = m.InsertNewTransaction(ctx, trx)
	if err != nil {
		panic(err)
	}
}
