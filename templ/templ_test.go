package templ

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTemplate_Render(t *testing.T) {
	type args struct {
		data any
	}
	tests := []struct {
		name         string
		templatePath string
		args         args
		wantOutPut   string
		wantErr      bool
	}{
		{
			name:         "hello template",
			templatePath: "./test_data/hello.templ",
			args: args{
				data: map[string]string{
					"Name": "goplop",
				},
			},
			wantOutPut: "Hello goplop",
		},
		{
			name:         "if template",
			templatePath: "./test_data/controlFlowIf.templ",
			args: args{
				data: map[string]any{
					"Print": true,
				},
			},
			wantOutPut: "Print If control",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templFile, err := os.Open(tt.templatePath)
			if err != nil {
				t.Fatal(err)
			}
			defer templFile.Close()
			templ, err := NewTemplate(templFile, templFile.Name())
			if err != nil {
				t.Fatal(err)
			}
			outPut := &bytes.Buffer{}
			if err := templ.Render(outPut, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Template.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutPut := outPut.String(); strings.TrimSpace(gotOutPut) != tt.wantOutPut {
				t.Errorf("Template.Render() = %v, want %v", gotOutPut, tt.wantOutPut)
			}
		})
	}
}
