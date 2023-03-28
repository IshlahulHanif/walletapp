package api

import (
	"Julo/walletapp/usecase/wallet"
	"Julo/walletapp/utils"
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
		walletUsecase, err := wallet.GetInstance(c)
		if err != nil {
			errFinal = err
			return
		}

		m = Module{
			usecase: usecase{
				wallet: walletUsecase,
			},
		}
	})

	return m, errFinal
}
