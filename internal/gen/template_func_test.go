package gen

import (
	"html/template"
	"os"
	"testing"
)

func TestAddTemplateFunc(t *testing.T) {
	tpl := "This is {{AddTest}}"

	AddTemplateFunc(template.FuncMap{
		"AddTest": func() string {
			return "AddTest return"
		},
	})

	// 解析模板
	tmpl, _ := template.New("test").Funcs(GetTemplateFuncList()).Parse(tpl)

	In := map[string]interface{}{
		"Tables": []string{"Test"},
		"Srv":    "TestSrv",
	}

	err := tmpl.Execute(os.Stdout, In)
	if err != nil {
		t.Log("err:", err.Error())
	}
}
