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
	Key        string // example: PRI 主键
	Default    string
	Extra      string
	Privileges string
	Comment    string
}

type IndexType int

const (
	IndexKey IndexType = iota
	PrimaryIndexKey
	UniqueIndexKey
)

type SqlIndex struct {
	Type        IndexType
	KeyName     string
	Unique      bool
	SeqInIndex  int
	ColumnName  []string
	Collation   string
	Cardinality string
	//Null         string
	IndexType    string
	Comment      string
	IndexComment string
	Expression   string
	Visible      string
}

type TableInfo struct {
	Name         string
	Columns      []*Column
	CreateTable  string
	PrimaryIndex *SqlIndex
	SqlIndexes   []*SqlIndex
}

// gen
type GenItem struct {
	Type    string
	Name    string
	Comment string
}

type ProtoItem struct {
	GenItem
}

type ProtoIndexItem struct {
	GenItem
	Fields     []string
	IndexItems []*ProtoItem
}

type ProtoContent struct {
	ProtoItems       []*ProtoItem
	PrimaryIndexItem *ProtoIndexItem
	IndexItem        []*ProtoIndexItem
}

type GoStructItem struct {
	GenItem
	Column *Column
}

type GoStructContent struct {
	GoStructItems []*GoStructItem
}

type PrimaryIndexItem struct {
	ProtoItem    *ProtoItem
	GoStructItem *GoStructItem
}

type Content struct {
	Srv              string
	TableInfo        *TableInfo
	ProtoContent     *ProtoContent
	GoStructContent  *GoStructContent
	Package          string
	GoPackage        string
	PrimaryIndexItem *PrimaryIndexItem
}

// regex
