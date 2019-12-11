package templater

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
)

type AppTemplate struct {
	Name  string
	Bg    string
	Tmpls []*Templater
}

type Templater struct {
	Tmpl string
	Dir  string
	Name string
}

var toFuncs = template.FuncMap{
	"ToCamel":      strcase.ToCamel,
	"ToLowerCamel": strcase.ToLowerCamel,
}

func (t *Templater) Create(curDir string, data interface{}) error {
	tmpl := template.Must(template.New(t.Name).Funcs(toFuncs).Parse(t.Tmpl))
	if err := os.MkdirAll(filepath.Join(curDir, t.Dir), 0755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(curDir, t.Dir, t.Name))
	if err != nil {
		return err
	}
	defer f.Close()
	if err := tmpl.Execute(f, data); err != nil {
		return err
	}
	return nil
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
