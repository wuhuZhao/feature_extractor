package parse

import (
	"errors"
	"strconv"
	"strings"
)

/**
* versio=1,tf_vector_len=100,type=sparse;
name=tmpName,op=random("jsonPath"),type=int64,index=1;
name=tmpName1,op=random("jsonPath1","jsonPath2"),type=string,index=2;
**/
type Parse struct {
	version    int
	tensorLen  int
	tensorType string
	root       *GraphNode
}

func (p *Parse) Parse(config string) error {
	line := strings.Split(config, ";")
	checkLine := []string{}
	for i := 0; i < len(line); i++ {
		if strings.HasPrefix(line[i], "#") {
			continue
		}
		checkLine = append(checkLine, line[i])
	}
	if err := p.check(checkLine); err != nil {
		return err
	}
	nodes := []*OperatorNode{}
	for i := 1; i < len(checkLine); i++ {
		node, err := p.parseLine(checkLine[i])
		if err != nil {
			return err
		}
		nodes = append(nodes, node)
	}
	return nil
}

// check: 检查第一行的元信息
func (p *Parse) check(line []string) error {
	if len(line) == 0 {
		return errors.New("line1 should contain version,tf_vector_len and type")
	}
	ll := strings.Split(line[0], ",")
	if len(ll) != 3 {
		return errors.New("line1 should contain version, tf_vector_len and type")
	}
	for i := 0; i < len(ll); i++ {
		if strings.HasPrefix(ll[i], "version") {
			v := strings.Split(ll[i], "=")
			if len(v) != 2 {
				return errors.New("line1 version should contain a number like version=1")
			}
			cv, err := strconv.Atoi(v[1])
			if err != nil {
				return errors.New("line1 version should contain a number like version=1")
			}
			p.version = cv
		} else if strings.HasPrefix(ll[i], "tf_vector_len") {
			t := strings.Split(ll[i], "=")
			if len(t) != 2 {
				return errors.New("line1 tf_vector_len should contain a number like tf_vector_len=100")
			}
			ct, err := strconv.Atoi(t[1])
			if err != nil {
				return errors.New("line1 tf_vector_len should contain a number like tf_vector_len=100")
			}
			p.tensorLen = ct
		} else if strings.HasPrefix(ll[i], "type") {
			t := strings.Split(ll[i], "=")
			if len(t) != 2 {
				return errors.New("line1 tf_vector_len should contain a number like type=sparse")
			}
			p.tensorType = t[1]
		}
	}
	if p.version == 0 || p.tensorLen == 0 || p.tensorType == "" {
		return errors.New("line1 should contain version,tf_vector_len and type")
	}
	return nil
}

// parseLine: 解析每一行使用的算子和输入输出和具体对应的索引
// name=tmpName,op=random(jsonPath),type=int64,index=1;
func (p *Parse) parseLine(line string) (*OperatorNode, error) {
	ll := strings.Split(line, ",")
	if len(ll) != 4 {
		return nil, errors.New("line shouble like name=tmpName,op=random(\"jsonPath\"),type=int64,index=1;")
	}
	res := &OperatorNode{}
	for i := 0; i < len(ll); i++ {
		if strings.HasPrefix(ll[i], "name") {
			n := strings.Split(ll[i], "=")
			if len(n) != 2 {
				return nil, errors.New("line should be like name=tmpName")
			}
			res.name = n[1]
		} else if strings.HasPrefix(ll[i], "op") {
			n := strings.Split(ll[i], "=")
			if len(n) != 2 {
				return nil, errors.New("line should be like op=random(jsonPath)")
			}
			funcIdx := strings.Index(n[1], "(")
			if funcIdx == -1 {
				return nil, errors.New("line should be like op=random(jsonPath)")
			}
			res.funcName = n[1][:funcIdx]
			if funcIdx+1 > len(n[1])-1 {
				return nil, errors.New("line shouble like op=random(jsonPath)")
			}
			paramsLine := n[1][funcIdx+1 : len(n[1])-1]
			if len(paramsLine) == 0 {
				continue
			}
			params := strings.Split(paramsLine, ",")
			res.params = params
		} else if strings.HasPrefix(ll[i], "type") {
			n := strings.Split(ll[i], "=")
			if len(n) != 2 {
				return nil, errors.New("line should be type=int64")
			}
			res.outType = n[1]
		} else if strings.HasPrefix(ll[i], "index") {
			n := strings.Split(ll[i], "=")
			if len(n) != 2 {
				return nil, errors.New("line should be index=1")
			}
			idx, err := strconv.Atoi(n[1])
			if err != nil {
				return nil, errors.New("line should be index=1, index is number")
			}
			res.index = int64(idx)
		}
	}
	return res, nil
}

func (p *Parse) generateGraph(op *OperatorNode) {

}
