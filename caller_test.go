package grpcdiff

import (
	"testing"
)

func TestCaller(t *testing.T) {
	c, err := NewCaller(":50051", ":50052", "pkg.DynamicPage", "Handle")
	if err != nil {
		t.Error(err)
	}

	r, err := c.Call(`{"content":"hello"}`)
	if err != nil {
		t.Error(err)
	}

	t.Error(r.Report())
	r, err = c.Call(`{"content":"h"}`)
	if err != nil {
		t.Error(err)
	}

	t.Error(r.Report())
}
