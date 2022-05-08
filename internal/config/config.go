package config

type SqlConfig struct {
	SqlDsn    string
	TableName string
}

type GenConfig struct {
	SqlConfigs []*SqlConfig
	SrvName    string
	SavePath   string
	Debug      bool
}
