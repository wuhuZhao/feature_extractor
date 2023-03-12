package parse

type OperatorNode struct {
	params   []string
	index    int64
	outType  string
	funcName string
	name     string
}

type GraphNode struct {
	meta     *OperatorNode
	name     string
	funcName string
	next     map[string]*GraphNode
	prev     map[string]*GraphNode
	end      bool
}

func NewGraph(config []*OperatorNode) *GraphNode {
	root := &GraphNode{}
	root.init(config)
	return root
}

func (g *GraphNode) init(config []*OperatorNode) {
	var 
}
