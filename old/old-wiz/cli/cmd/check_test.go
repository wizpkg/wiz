package cmd

import (
	"testing"
)

func TestBackend(t *testing.T) {
	testcases := map[string]string{
		"test.json":     "json",
		"test.prototxt": "prototxt",
	}
	for k, v := range testcases {
		t.Log(k, v)
		res, err := backend(k)
		t.Log(res, err)
		if err != nil {
			t.Fail()
		}
		if res != v {
			t.Fail()
		}
	}
}
