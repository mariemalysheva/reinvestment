package service

import "errors"

var (
	ErrNotAuthorized      = errors.New("not authorized")
	ErrAccountNotFound    = errors.New("account not found")
	ErrIncorrectSignature = errors.New("incorrect signature")
)
