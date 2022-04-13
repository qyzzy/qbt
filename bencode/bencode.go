package bencode

import (
	"bufio"
	"errors"
	"io"
)

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
	BENINT  BenType = 0x02
	BENLIST BenType = 0x03
	BENDICT BenType = 0x04
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

func (b *BenObject) Bencode(w io.Writer) int {
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	wLen := 0
	switch b.type_ {
	case BENSTR:
		str, _ := b.Str()
		wLen += EncodeString(bw, str)
	case BENINT:
		val, _ := b.Int()
		wLen += EncodeInt(bw, val)
	case BENLIST:
		bw.WriteByte('l')
		list, _ := b.List()
		for _, elem := range list {
			wLen += elem.Bencode(bw)
		}
		bw.WriteByte('e')
		wLen += 2
	case BENDICT:
		bw.WriteByte('d')
		dict, _ := b.Dict()
		for k, v := range dict {
			wLen += EncodeString(bw, k)
			wLen += v.Bencode(bw)
		}
		bw.WriteByte('e')
		wLen += 2
	}
	bw.Flush()
	return wLen
}

func EncodeString(w *bufio.Writer, val string) int {
	strLen := len(val)
	bw := bufio.NewWriter(w)
	wLen := writeDecimal(bw, strLen)
	bw.WriteByte(':')
	wLen++
	bw.WriteString(val)
	wLen += strLen
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

func EncodeInt(w *bufio.Writer, val int) int {
	bw := bufio.NewWriter(w)
	wLen := 0
	bw.WriteByte('i')
	wLen++
	nLen := writeDecimal(bw, val)
	wLen += nLen
	bw.WriteByte('e')
	wLen++
	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wLen
}

func writeDecimal(w *bufio.Writer, val int) int {
	len := 0
	if val == 0 {
		w.WriteByte('0')
		return 1
	}
	if val < 0 {
		w.WriteByte('-')
		len++
		val *= -1
	}

	return len
}
