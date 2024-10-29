package config

type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}

type snowflake struct {
	WorkerID      int64 `mapstructure:"worker-id"`
	DatancenterID int64 `mapstructure:"datancenter-id"`
}

type mySQL struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type redis struct {
	Addr     string
	Password string
}

type rabbitMQ struct {
	Addr     string
	Username string
	Password string
}

type elasticsearch struct {
	Addr string
	Host string
}

type config struct {
	Snowflake     snowflake
	MySQL         mySQL
	Redis         redis
	RabbitMQ      rabbitMQ
	Elasticsearch elasticsearch
}
