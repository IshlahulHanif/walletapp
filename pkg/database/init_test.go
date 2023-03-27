package database

import (
	"fmt"
	"testing"
)

func TestConnectDB(t *testing.T) {
	m, err := GetInstance(ConfigDatabase{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Dbname:   "wallet_db",
	})
	if err != nil {
		fmt.Println(err)
	}
	err = m.CloseDbConn()
	if err != nil {
		fmt.Println(err)
	}
}
