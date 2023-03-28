package wallet

import (
	"Julo/walletapp/entity"
	"Julo/walletapp/pkg/auth"
	"context"
	"github.com/IshlahulHanif/logtrace"
	"math"
	"time"
)

func (m Module) InitAccountWallet(ctx context.Context, customerID string) (string, error) {
	var (
		err error
	)

	// create token
	token := auth.GenerateTokenForUser(customerID)

	// TODO: should use tx.Commit() to avoid partially correct data
	// store user token
	err = m.repo.user.InsertUserProviderToken(ctx, customerID, token)
	if err != nil { // TODO: should I use upsert instead?
		return "", err
	}

	// create new wallet
	err = m.repo.wallet.CreateNewWallet(ctx, customerID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m Module) ChangeWalletStatus(ctx context.Context, token string, isEnable bool) (entity.Wallet, error) {
	var (
		err  error
		resp entity.Wallet
	)

	// get the user who hold this token
	userProvider, err := m.repo.user.GetUserProviderByToken(ctx, token)
	if err != nil {
		return resp, err
	}

	// TODO: improve by get the wallet via RETURNING clause in the query, or use Where is_enabled != :isEnable
	resp, err = m.repo.wallet.GetWalletByCustomerID(ctx, userProvider.CustomerID)
	if err != nil {
		return resp, err
	}

	// check if this wallet is already enabled/disabled
	if resp.IsEnabled == isEnable {
		if isEnable {
			err = entity.ConstErrAlreadyEnabled
		} else {
			err = entity.ConstErrAlreadyDisabled
		}
		logtrace.PrintLogErrorTrace(err)
		return entity.Wallet{}, err
	}

	// update the wallet status
	err = m.repo.wallet.UpdateWalletStatusByCustomerID(ctx, userProvider.CustomerID, isEnable)
	if err != nil {
		return resp, err
	}

	// TODO: should not do this if we use proper get & update
	resp.IsEnabled = isEnable

	return resp, nil
}

func (m Module) CheckWalletBalance(ctx context.Context, token string) (float64, error) {
	var (
		err error
	)

	// get the user who hold this token
	userProvider, err := m.repo.user.GetUserProviderByToken(ctx, token)
	if err != nil {
		return 0, err
	}

	wlt, err := m.repo.wallet.GetWalletByCustomerID(ctx, userProvider.CustomerID)
	if err != nil {
		return 0, err
	}

	if !wlt.IsEnabled {
		err = entity.ConstErrWalletDisabled
		logtrace.PrintLogErrorTrace(err)
		return 0, err
	}

	return wlt.Balance, nil
}

func (m Module) CheckWalletTransactionHistory(ctx context.Context, token string) ([]entity.TransactionHistory, error) {
	var (
		resp []entity.TransactionHistory
		err  error
	)

	// get the user who hold this token
	userProvider, err := m.repo.user.GetUserProviderByToken(ctx, token)
	if err != nil {
		return resp, err
	}

	// TODO: can optimize
	wlt, err := m.repo.wallet.GetWalletByCustomerID(ctx, userProvider.CustomerID)
	if err != nil {
		return resp, err
	}

	if !wlt.IsEnabled {
		err = entity.ConstErrWalletDisabled
		logtrace.PrintLogErrorTrace(err)
		return resp, err
	}

	resp, err = m.repo.transaction.GetAllTransactionByWalletID(ctx, wlt.ID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m Module) DepositMoneyToWallet(ctx context.Context, req UpdateWalletBalanceReq) (entity.Wallet, error) {
	var (
		resp entity.Wallet
		err  error
	)

	if req.Amount < 0 {
		err = entity.ConstErrInvalidAmount
		logtrace.PrintLogErrorTrace(err)
		return resp, err
	}

	return m.updateWalletBalance(ctx, req)
}

func (m Module) WithdrawMoneyFromWallet(ctx context.Context, req UpdateWalletBalanceReq) (entity.Wallet, error) {
	var (
		resp entity.Wallet
		err  error
	)

	if req.Amount > 0 {
		err = entity.ConstErrInvalidAmount
		logtrace.PrintLogErrorTrace(err)
		return resp, err
	}

	return m.updateWalletBalance(ctx, req)
}

func (m Module) updateWalletBalance(ctx context.Context, req UpdateWalletBalanceReq) (entity.Wallet, error) {
	var (
		resp       entity.Wallet
		err, errDB error
	)

	// TODO: really should use tx.Commit & tx.Rollback here

	defer func() {
		var (
			status  = ConstStatusSuccess
			trxType string
		)
		// if its error db, don't insert history
		if errDB != nil {
			return
		}

		// store history data
		if err != nil {
			status = ConstStatusFailed
		}

		if req.Amount < 0 {
			trxType = ConstTypeWithdrawal
		} else {
			trxType = ConstTypeDeposit
		}

		errInsertHistory := m.repo.transaction.InsertNewTransaction(ctx, entity.TransactionHistory{
			Status:       status,
			WalletID:     resp.ID, // TODO: what if we don't have the wallet ID bcs get wallet is failed?
			TransactedAt: time.Now(),
			Type:         trxType,
			Amount:       math.Abs(req.Amount),
			ReferenceID:  req.ReferenceID,
		})
		if errInsertHistory != nil {
			logtrace.PrintLogErrorTrace(errInsertHistory)
			// should do a rollback here
		}
	}()

	// get the user who hold this token
	userProvider, errDB := m.repo.user.GetUserProviderByToken(ctx, req.Token)
	if errDB != nil {
		return resp, errDB
	}

	resp, errDB = m.repo.wallet.GetWalletByCustomerID(ctx, userProvider.CustomerID)
	if errDB != nil {
		return resp, errDB
	}

	if !resp.IsEnabled {
		err = entity.ConstErrWalletDisabled
		logtrace.PrintLogErrorTrace(err)
		return entity.Wallet{}, err
	}

	// check if withdraw is possible
	if req.Amount < 0 && resp.Balance+req.Amount < 0 {
		err = entity.ConstErrNotEnoughMoney
		logtrace.PrintLogErrorTrace(err)
		return entity.Wallet{}, err
	}

	errDB = m.repo.wallet.IncrWalletAmount(ctx, userProvider.CustomerID, req.Amount)
	if errDB != nil {
		return entity.Wallet{}, errDB
	}

	// TODO: CAN REALLY OPTIMIZE, WHY YOU GET WALLET TWICE TO DB???
	resp, errDB = m.repo.wallet.GetWalletByCustomerID(ctx, userProvider.CustomerID)
	if errDB != nil {
		return resp, errDB
	}

	return resp, nil
}
