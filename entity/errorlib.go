package entity

import "errors"

var (
	ConstErrAlreadyEnabled  = errors.New("Already enabled")
	ConstErrAlreadyDisabled = errors.New("Already disabled")
	ConstErrNoRowsAffected  = errors.New("no rows affected")
	ConstErrWalletDisabled  = errors.New("Wallet disabled")
	ConstErrInvalidAmount   = errors.New("invalid amount")
	ConstErrNotEnoughMoney  = errors.New("not enough money")
)
