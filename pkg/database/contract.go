package database

type MethodDatabase interface {
	CloseDbConn() error
}
