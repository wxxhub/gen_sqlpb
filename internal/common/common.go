package common

const DefaultProtoFileName = "default_proto"
const DefaultStructFileName = "default_struct"

// flag
// --dsn "root:123456@tcp(www.wxxhome.com:3306)/test?tableName=new_table" --dsn "root:123456@tcp(www.wxxhome.com:3306)/test?tableName=user&srvName=UserSrv" --savePath cache --debug true
type Option struct {
	SavePath            string   `long:"savePath" description:"protobuf save path" default:"./"`
	DSN                 []string `long:"dsn" description:"data source name"`
	Debug               bool     `long:"debug" description:"print debug info"`
	Temples             []string `long:"temples" short:"t" description:"gen temples"`
	NotUseDefaultTemple bool     `shot:"nud" description:"not use default temples"`
}

// config
type DbConfig struct {
	Dsn       string
	DataBase  string
	TableName string
	SrvName   string
}

type ServiceConfig struct {
	DbConfig       *DbConfig
	SrvName        string
	SavePath       string
	FileName       string
	SqlSavePath    string
	SqlFileName    string
	StructSavePath string
	StructFileName string
	Package        string
	GoPackage      string
}

type GlobalConfig struct {
	Services map[string]*ServiceConfig
	Debug    bool
	Option   Option
}

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
	FName       string
	UpperName   string
	CamelName   string
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
	Column    *Column
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
