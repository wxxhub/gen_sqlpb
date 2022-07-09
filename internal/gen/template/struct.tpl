package {{.TableInfo.Name}}

type {{.TableInfo.UpperName}} struct {
{{- range $item := .GoStructItems}}
    {{$item.Name}} {{$item.Type}} `gorm:"{{$item.Name}}" json:"{{$item.Name}}"`
{{- end}}
}