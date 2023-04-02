package hash

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/wuhuZhao/feature_extractor/pkg/operator"
	"hash/fnv"
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
	hashStr := hex.EncodeToString(m.Sum(nil))
	h := fnv.New32a()
	h.Write([]byte(hashStr))
	return h.Sum32(), nil
}

func (s *StringToMd5Operator) Register() string {
	return "md5"
}
