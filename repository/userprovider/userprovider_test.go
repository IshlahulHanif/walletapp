package userprovider

import (
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"context"
	"fmt"
	"testing"
)

func TestModule_GetUserProviderByToken(t *testing.T) {
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

	err = m.UpsertUserProviderToken(ctx, "shlh-test", "new-token")
	if err != nil {
		panic(err)
	}
}

func TestModule_UpsertUserProviderToken(t *testing.T) {
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

	res, err := m.GetUserProviderByToken(ctx, "new-token")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
