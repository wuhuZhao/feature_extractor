package random

import "testing"

func TestRandomOperator(t *testing.T) {
	op := &RandomOperator{}
	if op.Register() != "random" {
		t.Fatalf("err: opName != %s\n", op.Register())
	}
	i, err := op.Handler(nil)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if data, ok := i.(int64); ok {
		t.Logf("get result: %v\n", data)
	} else {
		t.Fatal("handler error")
	}
}
