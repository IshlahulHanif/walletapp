package userprovider

import (
	"Julo/walletapp/entity"
	"Julo/walletapp/utils"
	"context"
	"database/sql"
	"errors"
	"github.com/IshlahulHanif/logtrace"
	"time"
)

func (m Module) UpsertUserProviderToken(ctx context.Context, customerID string, token string) error {
	var (
		err         error
		processTime = time.Now()
	)

	userProvider := entity.UserProvider{
		CustomerID: customerID,
		Token:      token,
		CreateTime: processTime,
		CreatedBy:  utils.ConstAppName,
		UpdateTime: processTime,
		UpdatedBy:  utils.ConstAppName,
	}

	_, err = m.database.NamedExec(ctx, ConstUpsertUserProviderToken, userProvider)
	if err != nil {
		logtrace.PrintLogErrorTrace(err)
		return err
	}

	return nil
}

func (m Module) GetUserProviderByToken(ctx context.Context, token string) (entity.UserProvider, error) {
	var (
		err          error
		userProvider entity.UserProvider
	)

	err = m.database.Get(ctx, &userProvider, ConstGetUserProviderByToken, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logtrace.PrintLogErrorTrace(err)
		return userProvider, err
	}

	return userProvider, nil
}
