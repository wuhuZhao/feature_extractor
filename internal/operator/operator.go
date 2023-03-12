package operator

type Operator interface {
	// 特定算子的调用逻辑
	Handler(in interface{}) (interface{}, error)
	// 算子注册到框架的名称
	Register() string
}

var _ error = (*TypeError)(nil)

type TypeError struct {
}

func (t *TypeError) Error() string {
	return "type is not match the operatorInput"
}
