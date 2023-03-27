package database

import (
	"context"
	"database/sql"
)

// TODO: use tx implementation to avoid error and multi commit
func (m Module) CloseDbConn() error {
	return m.dbConn.Close()
}

func (m Module) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return m.dbConn.NamedExecContext(ctx, query, arg)
}

func (m Module) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.dbConn.SelectContext(ctx, dest, query, args...)
}

func (m Module) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.dbConn.GetContext(ctx, dest, query, args...)
}

func (m Module) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return m.dbConn.ExecContext(ctx, query, args...)
}
