package bodygenerator

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	source := &SeperatedValueSource{
		value: []string{"1", "2", "3", "4"},
	}

	tmpl := `{"menuID":"{{.L1}}", "platformID": "5"}`

	s := &strings.Builder{}

	err := Generate(source, tmpl, s)
	if err != nil {
		t.Error(err)
	}
	want := `{"menuID":"1", "platformID": "5"}
{"menuID":"2", "platformID": "5"}
{"menuID":"3", "platformID": "5"}
{"menuID":"4", "platformID": "5"}
`

	if want != s.String() {
		t.Errorf("\nwant: %s\n got: %s", want, s)
	}
}
