package entity

import "errors"

var (
	ConstErrAlreadyEnabled  = errors.New("Already enabled")
	ConstErrAlreadyDisabled = errors.New("Already disabled")
	ConstErrNoRowsAffected  = errors.New("no rows affected")
)
