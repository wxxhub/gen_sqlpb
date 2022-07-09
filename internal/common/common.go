package common

// db
type Column struct {
	Field      string
	Type       string
	Collation  string
	Null       string
	Key        string
	Default    string
	Extra      string
	Privileges string
	Comment    string
}

type TableInfo struct {
	Name        string
	UpperName   string
	Columns     []*Column
	CreateTable string
}

// gen
type ProtoItem struct {
	Type      string
	Name      string
	NameUpper string
	Index     int
}

type GoStructItem struct {
	Type      string
	Name      string
	NameUpper string
}

type Table struct {
	Name      string
	UpperName string
}

type Content struct {
	Srv           string
	TableInfo     *TableInfo
	ProtoItems    []*ProtoItem
	GoStructItems []*GoStructItem
	Package       string
	GoPackage     string
}
