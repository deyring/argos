package models

import "errors"

// General Errors
var (
	ErrorNotImplemented = errors.New("not implemented")
)

// Transaction Errors
var (
	ErrorTransactionNameMissing   = errors.New("transaction name is missing")
	ErrorTransactionChecksMissing = errors.New("transaction checks are missing")
)

// EndpointCheck Errors
var (
	ErrorCheckNameMissing       = errors.New("check name is missing")
	ErrorCheckURLMissing        = errors.New("check url is missing")
	ErrorCheckURLInvalid        = errors.New("check url is invalid")
	ErrorCheckInvalidMethod     = errors.New("check method is invalid")
	ErrorCheckAssertionsMissing = errors.New("check assertions are missing")
)
