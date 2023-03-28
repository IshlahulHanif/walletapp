package api

import (
	"Julo/walletapp/entity"
	"Julo/walletapp/usecase/wallet"
	"context"
	"encoding/json"
	"errors"
	"github.com/IshlahulHanif/logtrace"
	"net/http"
	"strconv"
	"strings"
)

func (m Module) HandlerInitAccountWallet(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodPost {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	customerID := r.FormValue("customer_xid")

	token, err := m.usecase.wallet.InitAccountWallet(ctx, customerID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}

func (m Module) HandlerWallet(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		resp entity.JSend
	)

	switch r.Method {
	case http.MethodPost:
		m.HandlerEnableAccountWallet(w, r)
	case http.MethodGet:
		m.HandlerCheckWalletBalance(w, r)
	case http.MethodPatch:
		m.HandlerDisableAccountWallet(w, r)
	default:
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
	}
}

func (m Module) HandlerEnableAccountWallet(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodPost {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	token := strings.Trim(r.Header.Get("Authorization"), "Token ")

	wlt, err := m.usecase.wallet.ChangeWalletStatus(ctx, token, true)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			ID        string  `json:"id"`
			OwnedBy   string  `json:"owned_by"`
			Status    string  `json:"status"`
			EnabledAt string  `json:"enabled_at"`
			Balance   float64 `json:"balance"`
		}{
			ID:        wlt.ID,
			OwnedBy:   wlt.CustomerID,
			Status:    "enabled",
			EnabledAt: wlt.UpdateTime.Format("2006-01-02T15:04:05-07:00"),
			Balance:   wlt.Balance,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}

func (m Module) HandlerCheckWalletBalance(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodGet {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	token := strings.Trim(r.Header.Get("Authorization"), "Token ")

	wlt, err := m.usecase.wallet.CheckWalletBalance(ctx, token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	var status string
	if wlt.IsEnabled {
		status = "enabled"
	} else {
		status = "disabled"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			ID        string  `json:"id"`
			OwnedBy   string  `json:"owned_by"`
			Status    string  `json:"status"`
			EnabledAt string  `json:"enabled_at"`
			Balance   float64 `json:"balance"`
		}{
			ID:        wlt.ID,
			OwnedBy:   wlt.CustomerID,
			Status:    status,
			EnabledAt: wlt.UpdateTime.Format("2006-01-02T15:04:05-07:00"),
			Balance:   wlt.Balance,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}

func (m Module) HandlerCheckWalletTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodGet {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	token := strings.Trim(r.Header.Get("Authorization"), "Token ")

	trx, err := m.usecase.wallet.CheckWalletTransactionHistory(ctx, token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	var data []Transaction

	for _, tr := range trx {
		data = append(data, Transaction{
			ID:           tr.ID,
			Status:       tr.Status,
			TransactedAt: tr.TransactedAt.Format("2006-01-02T15:04:05-07:00"),
			Amount:       tr.Amount,
			RefID:        tr.ReferenceID,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: TransactionsHistory{
			Transaction: data,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}

func (m Module) HandlerDeposit(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodPost {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	token := strings.Trim(r.Header.Get("Authorization"), "Token ")

	refID := r.FormValue("reference_id")

	amountStr := r.FormValue("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || refID == "" || token == "" {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: "invalid request",
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	wlt, err := m.usecase.wallet.DepositMoneyToWallet(ctx, wallet.UpdateWalletBalanceReq{
		Token:       token,
		Amount:      amount,
		ReferenceID: refID,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	var status string
	if wlt.IsEnabled {
		status = "enabled"
	} else {
		status = "disabled"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			ID        string  `json:"id"`
			OwnedBy   string  `json:"owned_by"`
			Status    string  `json:"status"`
			EnabledAt string  `json:"enabled_at"`
			Balance   float64 `json:"balance"`
		}{
			ID:        wlt.ID,
			OwnedBy:   wlt.CustomerID,
			Status:    status,
			EnabledAt: wlt.UpdateTime.Format("2006-01-02T15:04:05-07:00"),
			Balance:   wlt.Balance,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}

func (m Module) HandlerWithdraw(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodPost {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	token := strings.Trim(r.Header.Get("Authorization"), "Token ")

	refID := r.FormValue("reference_id")

	amountStr := r.FormValue("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || refID == "" || token == "" {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: "invalid request",
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	wlt, err := m.usecase.wallet.WithdrawMoneyFromWallet(ctx, wallet.UpdateWalletBalanceReq{
		Token:       token,
		Amount:      amount,
		ReferenceID: refID,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	var status string
	if wlt.IsEnabled {
		status = "enabled"
	} else {
		status = "disabled"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			ID        string  `json:"id"`
			OwnedBy   string  `json:"owned_by"`
			Status    string  `json:"status"`
			EnabledAt string  `json:"enabled_at"`
			Balance   float64 `json:"balance"`
		}{
			ID:        wlt.ID,
			OwnedBy:   wlt.CustomerID,
			Status:    status,
			EnabledAt: wlt.UpdateTime.Format("2006-01-02T15:04:05-07:00"),
			Balance:   wlt.Balance,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}

func (m Module) HandlerDisableAccountWallet(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = context.Background()
		err  error
		resp entity.JSend
	)

	if r.Method != http.MethodPatch {
		err = errors.New("unsupported HTTP method")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		logtrace.PrintLogErrorTrace(err)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	token := strings.Trim(r.Header.Get("Authorization"), "Token ")

	wlt, err := m.usecase.wallet.ChangeWalletStatus(ctx, token, false)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logtrace.PrintLogErrorTrace(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			ID        string  `json:"id"`
			OwnedBy   string  `json:"owned_by"`
			Status    string  `json:"status"`
			EnabledAt string  `json:"enabled_at"`
			Balance   float64 `json:"balance"`
		}{
			ID:        wlt.ID,
			OwnedBy:   wlt.CustomerID,
			Status:    "enabled",
			EnabledAt: wlt.UpdateTime.Format("2006-01-02T15:04:05-07:00"),
			Balance:   wlt.Balance,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
	}
}
