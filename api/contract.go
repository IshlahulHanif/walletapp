package api

import "net/http"

type API interface {
	HandlerInitAccountWallet(w http.ResponseWriter, r *http.Request)
	HandlerWallet(w http.ResponseWriter, r *http.Request)
	HandlerEnableAccountWallet(w http.ResponseWriter, r *http.Request)
	HandlerCheckWalletBalance(w http.ResponseWriter, r *http.Request)
	HandlerCheckWalletTransactions(w http.ResponseWriter, r *http.Request)
	HandlerDeposit(w http.ResponseWriter, r *http.Request)
	HandlerWithdraw(w http.ResponseWriter, r *http.Request)
	HandlerDisableAccountWallet(w http.ResponseWriter, r *http.Request)
}
