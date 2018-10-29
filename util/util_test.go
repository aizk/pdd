package util

import "testing"

func TestTransformPids(t *testing.T) {
	s := TransformPids([]string{"1"})
	t.Log(s)
	if s != `["1"]` {
		t.Error("error")
	}
}
