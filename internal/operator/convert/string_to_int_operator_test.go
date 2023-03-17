package convert

import "testing"

func TestStringToIntOperator(t *testing.T) {
	op := &StringToIntOperator{}
	if op.Register() != "StringToInt" {
		t.Fatalf("err: opName != %s\n", op.Register())
	}
	i, err := op.Handler("123123")
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if data, ok := i.(int); ok {
		if i != 123123 {
			t.Fatal("handler error")
		}
		t.Logf("get result: %v\n", data)
	} else {
		t.Fatal("handler error")
	}
}
