package templ_test

import (
	"bytes"
	"log"
	"os"

	"github.com/ProImpact/goplop/templ"
)

func ExampleTemplate() {
	templateSrc := `Hello {{ .User }}`
	buf := &bytes.Buffer{}
	_, err := buf.WriteString(templateSrc)
	if err != nil {
		log.Fatal(err)
	}
	templ, err := templ.NewTemplate(buf, "exapmle")
	if err != nil {
		log.Fatal(err)
	}
	templ.Render(os.Stdout, map[string]string{
		"User": "goplop",
	})
	// Output:
	// Hello goplop
}
