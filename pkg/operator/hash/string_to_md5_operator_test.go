package hash

import "testing"

func TestStringToMd5Operator(t *testing.T) {
	op := &StringToMd5Operator{}
	if op.Register() != "md5" {
		t.Fatalf("err: opName != %s\n", op.Register())
	}
	i, err := op.Handler("testfortest")
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if data, ok := i.(uint32); ok {
		t.Logf("get result: %v\n", data)
	} else {
		t.Fatal("handler error")
	}
}
