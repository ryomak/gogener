package templater

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
)

type Tmpl interface {
	String() (string, error)
}

// AppTemplate is info for creating app
type AppTemplate struct {
	Name  string
	Bg    string
	Tmpls []*Templater
}

// Templater is stuct for executing template
type Templater struct {
	Tmpl Tmpl
	Dir  string
	Name string
}

var toFuncs = template.FuncMap{
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}

// Create is creating file with template (using go fmt...)
func (t *Templater) Create(curDir string, data interface{}) error {
	tmplStr, err := t.Tmpl.String()
	if err != nil {
		return err
	}
	dir := filepath.Join(curDir, t.Dir)
	fullPath := filepath.Join(curDir, t.Dir, t.Name)
	tmpl := template.Must(template.New(t.Name).Delims("[[", "]]").Funcs(toFuncs).Parse(tmplStr))

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		return err
	}
	fmtBuf, err := imports.Process(dir, buf.Bytes(), nil)
	if err != nil {
		fmtBuf = buf.Bytes()
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return ioutil.WriteFile(fullPath, fmtBuf, 0644)
}
