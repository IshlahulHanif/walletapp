package userprovider

import (
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"github.com/IshlahulHanif/logtrace"
	"sync"
)

var (
	m    Module
	once sync.Once
)

func GetInstance(c utils.Config) (Module, error) {
	var (
		errFinal error
	)

	once.Do(func() {
		db, err := database.GetInstance(c.DatabaseConfig)
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		m = Module{
			database: db,
		}
	})

	return m, errFinal
}
