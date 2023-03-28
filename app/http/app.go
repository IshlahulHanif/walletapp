package main

import (
	"Julo/walletapp/api"
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"github.com/IshlahulHanif/logtrace"
	"net/http"
)

func main() {
	// init config
	conf := utils.Config{
		DatabaseConfig: database.ConfigDatabase{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			Dbname:   "wallet_db",
		},
	}

	// init logtrace settings

	logtrace.SetProjectName("walletapp")
	logtrace.SetIsPrintFromContentRoot(true)
	logtrace.SetIsPrintFunctionName(true)
	logtrace.SetIsPrintNewline(true)
	logtrace.SetIsUseTabSeparator(false)

	// init http api
	httpApi, err := api.GetInstance(conf)
	if err != nil {
		return
	}

	// register handlers
	http.HandleFunc("/api/v1/init", httpApi.HandlerInitAccountWallet)
	http.HandleFunc("/api/v1/wallet", httpApi.HandlerWallet)
	http.HandleFunc("/api/v1/wallet/transactions", httpApi.HandlerCheckWalletTransactions)
	http.HandleFunc("/api/v1/wallet/deposits", httpApi.HandlerDeposit)
	http.HandleFunc("/api/v1/wallet/withdrawals", httpApi.HandlerWithdraw)

	err = http.ListenAndServe("", nil)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return
	}
}
