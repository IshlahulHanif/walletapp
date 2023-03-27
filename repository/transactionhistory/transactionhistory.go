package transactionhistory

import (
	"Julo/walletapp/entity"
	"context"
	"database/sql"
	"errors"
	"github.com/IshlahulHanif/logtrace"
	"time"
)

func (m Module) InsertNewTransaction(ctx context.Context, trx entity.TransactionHistory) error {
	var (
		err         error
		processTime = time.Now()
	)

	trx.TransactedAt = processTime

	_, err = m.database.NamedExec(ctx, ConstInsertNewTransaction, trx)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return err
	}

	return nil
}

func (m Module) GetAllTransactionByWalletID(ctx context.Context, walletID int64) ([]entity.TransactionHistory, error) {
	var (
		err     error
		trxData []entity.TransactionHistory
	)

	err = m.database.Select(ctx, &trxData, ConstGetAllTransactionByWalletID, walletID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logtrace.PrintLogErrorTrace(err)
		return trxData, err
	}

	return trxData, nil
}
