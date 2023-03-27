package wallet

import (
	"Julo/walletapp/entity"
	"Julo/walletapp/pkg/auth"
	"context"
	"github.com/IshlahulHanif/logtrace"
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
