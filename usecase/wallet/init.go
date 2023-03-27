package wallet

import (
	"Julo/walletapp/repository/transactionhistory"
	"Julo/walletapp/repository/userprovider"
	"Julo/walletapp/repository/wallet"
	"Julo/walletapp/utils"
	"github.com/IshlahulHanif/logtrace"
	"sync"
)

var (
	m    Module
	once sync.Once
)

func GetInstance(c utils.Config) (Module, error) { //TODO: think if this should not be a singleton
	var (
		errFinal error
	)

	once.Do(func() {
		transactionRepo, err := transactionhistory.GetInstance(c)
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		userRepo, err := userprovider.GetInstance(c)
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		walletRepo, err := wallet.GetInstance(c)
		if err != nil {
			errFinal = err
			logtrace.PrintLogErrorTrace(err)
			return
		}

		m = Module{
			repo: repo{
				transaction: transactionRepo,
				user:        userRepo,
				wallet:      walletRepo,
			},
		}
	})

	return m, errFinal
}
