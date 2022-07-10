package {{.TableInfo.Name}}

type {{.TableInfo.CamelName}} struct {
{{- range $item := .GoStructContent.GoStructItems}}
    {{$item.Name}} {{$item.Type}} `gorm:"type:{{$item.Column.Type}}" json:"{{$item.Name}}" `
{{- end}}
}

func ({{.TableInfo.FName}} *{{.TableInfo.CamelName}}) TableName() string {
    return "{{.TableInfo.Name}}"
}