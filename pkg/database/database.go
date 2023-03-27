package database

func (m Module) CloseDbConn() error {
	return m.dbConn.Close()
}
