package database

import (
	"context"
	"database/sql"
)

type MethodDatabase interface {
	CloseDbConn() error
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
}
