package internal

import (
	"encoding/json"
	"errors"
	"fe_extractor/internal/operator"
	"fe_extractor/internal/operator/convert"
	"fe_extractor/internal/operator/hash"
	"fe_extractor/internal/operator/random"
	"fe_extractor/internal/parse"
	"fmt"
)

var operators map[string]operator.Operator

// 注册所有实现的算子
func init() {
	stringtoInt := &convert.StringToIntOperator{}
	stmo := &hash.StringToMd5Operator{}
	ro := &random.RandomOperator{}
	operators = make(map[string]operator.Operator)
	operators[stringtoInt.Register()] = stringtoInt
	operators[stmo.Register()] = stmo
	operators[ro.Register()] = ro
}

type Handler struct {
	p *parse.Parse
}

func NewHandler() *Handler {
	return &Handler{p: parse.NewParse()}
}

func (h *Handler) Init(config string) {
	h.p.Parse(config)
}

func (h *Handler) Handle(input string) ([]interface{}, error) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	output := make([]interface{}, h.p.TensorLen())
	h.p.Root.OP(func(node *parse.OperatorNode) interface{} {
		op, ok := operators[node.FuncName()]
		if !ok {
			panic(fmt.Errorf("have not op is :%s, please check your config", node.FuncName()))
		}
		var out interface{}
		// 区分一下  不想改数组了
		if len(node.Params()) == 1 {
			if v, ok := data[node.Params()[0]]; ok {
				out, err = op.Handler(v)
				if err != nil {
					panic(err)
				}
			} else if v, ok := result[node.Params()[0]]; ok {
				out, err = op.Handler(v)
				if err != nil {
					panic(err)
				}
			} else {
				panic(errors.New("can not get the params from result and input, please check your config or data"))
			}
		} else {
			ins := []interface{}{}
			for i := 0; i < len(node.Params()); i++ {
				if v, ok := data[node.Params()[i]]; ok {
					ins = append(ins, v)
				} else if v, ok := result[node.Params()[0]]; ok {
					ins = append(ins, v)
				} else {
					panic(errors.New("can not get the params from result and input, please check your config or data"))
				}
			}
			out, err = op.Handler(ins)
			if err != nil {
				panic(err)
			}
		}
		result[node.Name()] = out
		if node.Index() != -1 {
			output[node.Index()] = out
		}
		return out
	})

	return output, nil

}
