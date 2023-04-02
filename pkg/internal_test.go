package pkg

import (
	"io/ioutil"
	"testing"
)

func TestHandler(t *testing.T) {
	configBytes, err := ioutil.ReadFile("../fe.config")
	if err != nil {
		t.Fatalf("read config err: %v\n", err)
	}
	h := NewHandler()
	h.Init(string(configBytes))
	json, err := ioutil.ReadFile("../test.json")
	if err != nil {
		t.Fatalf("get json error: %v\n", err)
	}
	i, err := h.Handle(string(json))
	if err != nil {
		t.Fatalf("convert err: %v\n", err)
	}
	t.Logf("get data: %v\n", i)
}
