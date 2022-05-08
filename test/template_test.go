package test

import (
	"os"
	"testing"
	"text/template"

	_ "embed"
)

//go:embed ../internal/gen/template/proto.tpl
var protoTpl string

func TestTemplate(t *testing.T) {
	// 解析模板
	tmpl, _ := template.New("test").Parse(protoTpl)

	In := map[string]interface{}{
		"Tables": []string{"Test"},
		"Srv":    "TestSrv",
	}

	err := tmpl.Execute(os.Stdout, In)
	if err != nil {
		t.Log("err:", err.Error())
	}

}
