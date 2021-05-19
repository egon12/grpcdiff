package grpcdiff

import "testing"

func TestPrintDiff(t *testing.T) {
	a := `
1
2
3
4
5
`
	b := `
2
3
4
4
5
`

	PrintDiff([]byte(a), []byte(b))
}
