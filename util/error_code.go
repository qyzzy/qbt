package util

import "errors"

var (
	ErrorNumber         = errors.New("expect num")
	ErrorColon          = errors.New("expect colon")
	ErrorType           = errors.New("wrong type")
	ErrorInvalidBencode = errors.New("invalid bencode")
)
