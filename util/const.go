package util

import "errors"

var (
	ErrConflict                    = errors.New("error conflict")
	ErrInsert                      = errors.New("error inserting into db")
	ErrDelete                      = errors.New("error deleting from db")
	ErrInvalidToken                = errors.New("error invalid token")
	ErrRecordNotFound              = errors.New("record not found")
	ErrInvalidFieldName            = errors.New("error invalid field name")
	ErrNilCacheValue               = errors.New("error nil cache value")
	ErrUnexpectedType              = errors.New("error unexpected type")
	ErrMaxVerificationLimitReached = errors.New("max error verification limit reached")
	ErrExpiredToken                = errors.New("error expired token")
)

const (
	AppName = "NFTLookingGlass"
)
