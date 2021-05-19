package bodygenerator

import (
	"io"
	"text/template"
)

func Generate(source Source, tmplInput string, out io.Writer) error {
	tmpl, err := template.New("base").Parse(tmplInput + "\n")
	if err != nil {
		return err
	}

	for _, value := range source.GetValue() {
		tmpl.Execute(out, value)
	}

	return nil
}
