package config

type SqlConfig struct {
	SqlDsn    string
	TableName string
	SrvName   string
}

type ServiceConfig struct {
	SqlConfigs map[string]*SqlConfig
	SrvName    string
	SavePath   string
	FileName   string
	Package    string
	GoPackage  string
}

type GlobalConfig struct {
	Services map[string]*ServiceConfig
	Debug    bool
}
