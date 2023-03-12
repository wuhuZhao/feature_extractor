package random

import (
	"fe_extractor/internal/operator"
	"math/rand"
)

var _ operator.Operator = (*RandomOperator)(nil)

type RandomOperator struct {
}

func (r *RandomOperator) Handler(in interface{}) (interface{}, error) {
	return rand.Int63(), nil
}

func (r *RandomOperator) Register() string {
	return "random"
}
