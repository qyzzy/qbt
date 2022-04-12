package bencode

import "errors"

var (
	ErrorNumber         = errors.New("expect num")
	ErrorColon          = errors.New("expect colon")
	ErrorType           = errors.New("wrong type")
	ErrorInvalidBencode = errors.New("invalid bencode")
)

type BenType uint8

type BenValue interface{}

const (
	BENSTR  BenType = 0x01
	BENINT  BenType = 0x01
	BENLIST BenType = 0x01
	BENDICT BenType = 0x01
)

type BenObject struct {
	type_ BenType
	val   BenValue
}

func (b *BenObject) Str() (string, error) {
	if b.type_ != BENSTR {
		return "", ErrorType
	}
	return b.val.(string), nil
}

func (b *BenObject) Int() (int, error) {
	if b.type_ != BENINT {
		return 0, ErrorType
	}
	return b.val.(int), nil
}

func (b *BenObject) List() ([]*BenObject, error) {
	if b.type_ != BENLIST {
		return nil, ErrorType
	}
	return b.val.([]*BenObject), nil
}

func (b *BenObject) Dict() (map[string]*BenObject, error) {
	if b.type_ != BENDICT {
		return nil, ErrorType
	}
	return b.val.(map[string]*BenObject), nil
}
