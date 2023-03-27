package userprovider

import (
	"Julo/walletapp/entity"
	"context"
)

type Repository interface {
	UpsertUserProviderToken(ctx context.Context, customerID string, token string) error
	InsertUserProviderToken(ctx context.Context, customerID string, token string) error
	GetUserProviderByToken(ctx context.Context, token string) (entity.UserProvider, error)
}
