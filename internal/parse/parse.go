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
}

type OperatorNode struct {
	in      []string
	out     string
	index   int64
	outType string
}

func (p *Parse) Parse(config string) {
	line := strings.Split(config, ";")
	checkLine := []string{}
	for i := 0; i < len(line); i++ {
		if strings.HasPrefix(line[i], "#") {
			continue
		}
		checkLine = append(checkLine, line[i])
	}
	p.check(checkLine)

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
