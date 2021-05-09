package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func PrintDiff(a, b []byte) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(a), string(b), true)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
