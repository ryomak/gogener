package recipe

import (
	"io"
	"os"
	"testing"

	"github.com/ryomak/gogener/templater"
)

type testRecipe struct{}

func (*testRecipe) Do(*templater.AppTemplate) bool {
	return true
}
func TestNew(t *testing.T) {
	type input struct {
		out    io.Writer
		params Params
	}
	type TestCase struct {
		input  input
		expect Recipe
	}
	ts := []TestCase{
		{
			input: input{
				out: os.Stdout,
				params: Params{
					CurrentDir: "cur",
					AppName:    "app",
					ModName:    "mod",
				},
			},
			expect: &testRecipe{},
		},
	}
	New(ts[0].input.out, ts[0].input.params)
}

func TestDo(t *testing.T) {
	type testCase struct {
		input  *templater.AppTemplate
		expect bool
	}
	ts := []testCase{
		{
			input: &templater.AppTemplate{
				Name:  "name",
				Bg:    "bg",
				Tmpls: []*templater.Templater{},
			},
			expect: true,
		},
		{
			input: &templater.AppTemplate{
				Name:  "name",
				Bg:    "",
				Tmpls: []*templater.Templater{},
			},
			expect: true,
		},
		{
			input:  nil,
			expect: false,
		},
	}
	for i, v := range ts {
		recipe := New(os.Stdout, Params{
			AppName: "app",
		})
		if actual := recipe.Do(v.input); actual != v.expect {
			t.Fatalf("[%d]expect %v,actual %v ", i, actual, v.expect)
		}
	}
}
