package wallet

import (
	"Julo/walletapp/entity"
	"Julo/walletapp/utils"
	"context"
	"database/sql"
	"errors"
	"github.com/IshlahulHanif/logtrace"
	"time"
)

func (m Module) CreateNewWallet(ctx context.Context, customerID string) error {
	var (
		err         error
		processTime = time.Now()
	)

	wallet := entity.Wallet{
		CustomerID: customerID,
		Balance:    0,
		CreateTime: processTime,
		CreatedBy:  utils.ConstAppName,
	}

	_, err = m.database.NamedExec(ctx, ConstCreateNewWallet, wallet)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return err
	}

	return nil
}

func (m Module) UpdateWalletAmount(ctx context.Context, customerID string, amount float64) error { //TODO: use struct param?
	var (
		err         error
		processTime = time.Now()
	)

	wallet := entity.Wallet{
		CustomerID: customerID,
		Balance:    amount,
		UpdateTime: processTime,
		UpdatedBy:  utils.ConstAppName,
	}

	_, err = m.database.NamedExec(ctx, ConstUpdateWalletAmount, wallet)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return err
	}

	return nil
}

func (m Module) IncrWalletAmount(ctx context.Context, customerID string, amount float64) error { //TODO: use struct param?
	var (
		err         error
		processTime = time.Now()
	)

	wallet := entity.Wallet{
		CustomerID: customerID,
		Balance:    amount,
		UpdateTime: processTime,
		UpdatedBy:  utils.ConstAppName,
	}

	_, err = m.database.NamedExec(ctx, ConstIncrWalletAmount, wallet)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return err
	}

	return nil
}

func (m Module) GetWalletAmountByCustomerID(ctx context.Context, customerID string) (float64, error) {
	var (
		err    error
		amount float64
	)

	err = m.database.Get(ctx, &amount, ConstGetWalletAmountByCustomerID, customerID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logtrace.PrintLogErrorTrace(err)
		return amount, err
	}

	return amount, nil
}

func (m Module) UpdateWalletStatusByCustomerID(ctx context.Context, customerID string, isEnabled bool) error {
	var (
		err         error
		processTime = time.Now()
	)

	wallet := entity.Wallet{
		CustomerID: customerID,
		IsEnabled:  isEnabled,
		UpdateTime: processTime,
		UpdatedBy:  utils.ConstAppName,
	}

	_, err = m.database.NamedExec(ctx, ConstUpdateWalletStatusByCustomerID, wallet)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return err
	}

	return nil
}
