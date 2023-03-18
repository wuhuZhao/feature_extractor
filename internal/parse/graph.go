package parse

import "log"

type OperatorNode struct {
	params   []string
	index    int64
	outType  string
	funcName string
	name     string
}

func (op *OperatorNode) Params() []string {
	return op.params
}

func (op *OperatorNode) Index() int64 {
	return op.index
}

func (op *OperatorNode) OutType() string {
	return op.outType
}

func (op *OperatorNode) FuncName() string {
	return op.funcName
}

func (op *OperatorNode) Name() string {
	return op.name
}

type GraphNode struct {
	meta       *OperatorNode
	name       string
	next       map[string]*GraphNode
	dependence []string
	end        bool
}

// NewGraph: create graph
func NewGraph(config []*OperatorNode) *GraphNode {
	root := &GraphNode{name: "root"}
	root.init(config)
	return root
}

// init: init
func (g *GraphNode) init(config []*OperatorNode) {
	dig := map[string]int{}
	out := map[string]struct{}{}
	om := map[string]*OperatorNode{}
	state := map[string]*GraphNode{}
	//需要一个前向节点与后向节点的map
	pte := map[string][]string{}
	for i := 0; i < len(config); i++ {
		om[config[i].name] = config[i]
		dig[config[i].name] += len(config[i].params)
		for j := 0; j < len(config[i].params); j++ {
			pte[config[i].params[j]] = append(pte[config[i].params[j]], config[i].name)
		}
	}
	for k := range om {
		if _, ok := pte[k]; !ok {
			out[k] = struct{}{}
		}
	}
	for len(dig) > 0 {
		for k, v := range dig {
			if v != 0 {
				continue
			}
			node := &GraphNode{meta: om[k], name: k, dependence: om[k].params, next: map[string]*GraphNode{}, end: false}
			if _, ok := out[node.name]; ok {
				node.end = true
			}
			if len(om[k].params) == 0 {
				g.next[k] = node
			} else {
				for j := 0; j < len(om[k].params); j++ {
					state[om[k].params[j]].next[k] = node
				}

			}
			for j := 0; j < len(pte[k]); j++ {
				dig[pte[k][j]]--
			}
			state[k] = node
			delete(dig, k)
		}
	}
}

func (root *GraphNode) OP(handler func(node *OperatorNode) interface{}) error {
	defer func() {
		if p := recover(); p != nil {
			log.Fatalf("[OP] ERROR : %v\n", p)
		}
	}()
	cur := []*GraphNode{root}
	for len(cur) != 0 {
		next := []*GraphNode{}
		for i := 0; i < len(cur); i++ {
			if cur[i].name != "root" {
				handler(cur[i].meta)
			}
			for _, v := range cur[i].next {
				next = append(next, v)
			}
		}
		cur = next
	}
	return nil
}
