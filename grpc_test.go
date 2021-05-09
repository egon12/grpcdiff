package main

import (
	"testing"
)

func TestSimpleCall(t *testing.T) {
	b, err := SimpleCall("localhost:50051", "pkg.DynamicPage", "Handle1", `{"content":"hello"}`)

	t.Error(string(b))
	t.Error(err)
}
