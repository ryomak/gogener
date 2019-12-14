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

// AppTemplate is info for creating app
type AppTemplate struct {
	Name  string
	Bg    string
	Tmpls []*Templater
}

// Templater is stuct for executing template
type Templater struct {
	Tmpl string
	Dir  string
	Name string
}

var toFuncs = template.FuncMap{
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}

// Create is creating file with template (using go fmt...)
func (t *Templater) Create(curDir string, data interface{}) error {
	dir := filepath.Join(curDir, t.Dir)
	fullPath := filepath.Join(curDir, t.Dir, t.Name)
	tmpl := template.Must(template.New(t.Name).Funcs(toFuncs).Parse(t.Tmpl))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
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

/*
***
*** for template file
***
type Templater struct {
	SourcePath string
	OutPath    string
	Data       interface{}
}

func (t *Templater) Create() error {
	tmpl := template.Must(template.ParseFiles(t.SourcePath))
	if err := os.MkdirAll(filepath.Dir(t.OutPath), 0755); err != nil {
		return err
	}
	f, err := os.Create(t.OutPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := tmpl.Execute(f, t.Data); err != nil {
		return err
	}
	return nil
}
*/
