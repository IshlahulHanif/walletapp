package wallet

import (
	"Julo/walletapp/pkg/database"
	"Julo/walletapp/utils"
	"github.com/IshlahulHanif/logtrace"
	"sync"
)

// TODO: optimize singleton initiation, the singleton should be the Memory instance as it represents redis/postgres/other persistence with connection
var (
	m       Module
	mem     MemoryResource
	once    sync.Once
	onceMem sync.Once
)

func GetInstance(c utils.Config) (Module, error) {
	var (
		errFinal error
	)

	once.Do(func() {
		memory, err := GetMemory()
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		db, err := database.GetInstance(c.DatabaseConfig)
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		m = Module{
			memory:   memory,
			database: db,
		}
	})

	return m, errFinal
}

// TODO: actually use memcache
func GetMemory() (MemoryResource, error) {
	var (
		err error
	)

	onceMem.Do(func() {
		mem = MemoryResource{}
	})

	return mem, err
}
