package gen

import (
	"html/template"
	"os"
	"testing"
)

type MyMap map[string]string

func LastMapIndex(name string) string {
	if len(name) > 0 {
		return "Hello " + name
	}
	return "666"

}

func TestTemFunc(t *testing.T) {
	myMap := MyMap{}
	myMap["foo"] = "bar"

	tpl := template.New("template test")
	tpl = tpl.Funcs(template.FuncMap{"LastMapIndex": LastMapIndex})
	tpl = template.Must(tpl.Parse("<p>\nThe last index of this map is: {{.foo|LastMapIndex}}.\n</p>\n"))
	tpl.Execute(os.Stdout, myMap)
}
