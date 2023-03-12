package internal

import (
	"fe_extractor/internal/operator"
	"fe_extractor/internal/operator/convert"
	"fe_extractor/internal/operator/hash"
	"fe_extractor/internal/operator/random"
)

var operators map[string]operator.Operator

// 注册所有实现的算子
func init() {
	stringtoInt := &convert.StringToIntOperator{}
	stmo := &hash.StringToMd5Operator{}
	ro := &random.RandomOperator{}

	operators[stringtoInt.Register()] = stringtoInt
	operators[stmo.Register()] = stmo
	operators[ro.Register()] = ro
}
