package templater

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDo(t *testing.T) {
	type TestCase struct {
		t *Templater
		e string
	}
	ts := []TestCase{
		{
			t: &Templater{
				Tmpl: localTmpl{`aaa`},
				Dir:  "test",
				Name: "aaa.go",
			},
			e: "aaa",
		},
		{
			t: &Templater{
				Tmpl: localTmpl{`[[.Name]]`},
				Dir:  "test",
				Name: "aaa.go",
			},
			e: "name_example",
		},
		{
			t: &Templater{
				Tmpl: localTmpl{`[[ .Name |ToCamel]]`},
				Dir:  "test",
				Name: "aaa.go",
			},
			e: "NameExample",
		},
		{
			t: &Templater{
				Tmpl: localTmpl{`[[ .Name |ToLowerCamel]]`},
				Dir:  "test",
				Name: "aaa.go",
			},
			e: "nameExample",
		},
	}
	curDir := "cur"
	data := map[string]string{
		"Name": "name_example",
	}
	for _, v := range ts {
		if err := v.t.Create(curDir, data); err != nil {
			t.Fatalf("error:%v", err)
		}
		f, _ := os.Open(filepath.Join(curDir, v.t.Dir, v.t.Name))
		b, _ := ioutil.ReadAll(f)
		t.Log(filepath.Join(curDir, v.t.Dir))
		os.RemoveAll(filepath.Join(curDir))
		if string(b) != v.e {
			t.Fatalf("expected:%v,actual:%v", v.e, string(b))
		}
	}
}

func TestErrorDo(t *testing.T) {

	curDir := "cur"
	data := map[string]string{
		"Name": "name_example",
	}
	type TestCase struct {
		t *Templater
		e string
	}

	ts := TestCase{
		t: &Templater{
			Tmpl: localTmpl{`[[ .Name |ToLowerCamel]]`},
			Dir:  "test",
			Name: "aaa.go",
		},
		e: "nameExample",
	}
	os.Create(filepath.Join(curDir))
	if err := ts.t.Create(curDir, data); err == nil {
		os.RemoveAll(filepath.Join(curDir))
		t.Fatalf("expected error but nil")
	} else {
		os.RemoveAll(filepath.Join(curDir))
	}

	if err := ts.t.Create(curDir, nil); err == nil {
		os.RemoveAll(filepath.Join(curDir))
		t.Fatalf("expected error but nil")
	} else {
		os.RemoveAll(filepath.Join(curDir))
	}
}

type localTmpl struct {
	tmplStr string
}

func (s localTmpl) String() (string, error) {
	// ceternly error nil because local str
	return s.tmplStr, nil
}
