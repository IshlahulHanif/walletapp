package database

import (
	"fmt"
	"github.com/IshlahulHanif/logtrace"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
)

var (
	m    Module
	once sync.Once
)

func GetInstance(c ConfigDatabase) (MethodDatabase, error) { //TODO: can make improvement for multiple db connection
	var (
		errFinal error
	)

	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Dbname)

		// Open a connection to the database
		dbConn, err := sqlx.Open("postgres", psqlInfo)
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		// Test connection
		err = dbConn.Ping()
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		m = Module{
			dbConn: dbConn,
		}
	})

	return m, errFinal
}
