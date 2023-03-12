package convert

import (
	"fe_extractor/internal/operator"
	"strconv"
)

var _ operator.Operator = (*StringToIntOperator)(nil)

// 将string转换成int
type StringToIntOperator struct {
}

func (s *StringToIntOperator) Handler(in interface{}) (interface{}, error) {
	cur, ok := in.(string)
	if !ok {
		return nil, &operator.TypeError{}
	}
	result, err := strconv.Atoi(cur)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *StringToIntOperator) Register() string {
	return "StringToInt"
}
