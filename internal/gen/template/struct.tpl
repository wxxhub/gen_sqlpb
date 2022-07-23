package {{.TableInfo.Name}}

type {{.TableInfo.Name|StringCamel}} struct {
{{- range $item := .GoStructContent.GoStructItems}}
    {{$item.Name|StringCamel}} {{$item.Type}} `gorm:"type:{{$item.Column.Type}}" json:"{{$item.Name}}"` {{$item.Comment|AddNote}}
{{- end}}
}

func ({{.TableInfo.Name|StringLowFirst}} *{{.TableInfo.Name|StringCamel}}) TableName() string {
    return "{{.TableInfo.Name}}"
}