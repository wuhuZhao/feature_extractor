package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fe_extractor/internal/operator"
	"unsafe"
)

var _ operator.Operator = (*StringToMd5Operator)(nil)

type StringToMd5Operator struct {
}

func (s *StringToMd5Operator) Handler(in interface{}) (interface{}, error) {
	cur, ok := in.(string)
	if !ok {
		return nil, &operator.TypeError{}
	}
	m := md5.New()
	m.Write(*((*[]byte)(unsafe.Pointer(&cur))))
	return hex.EncodeToString(m.Sum(nil)), nil
}

func (s *StringToMd5Operator) Register() string {
	return "md5"
}
