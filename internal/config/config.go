package config

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
}
