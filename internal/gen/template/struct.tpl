package {{.TableInfo.Name}}

type {{.TableInfo.CamelName}} struct {
{{- range $item := .GoStructItems}}
    {{$item.Name}} {{$item.Type}} `gorm:"{{$item.Name}}" json:"{{$item.Name}}"`
{{- end}}
}

func ({{.TableInfo.FName}} *{{.TableInfo.CamelName}}) TableName() string {
    return "{{.TableInfo.Name}}"
}