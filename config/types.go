package config

type EtcdConf struct {
	Addr string
}

type AsynqRedisConf struct {
	Addr string
	Pass string
	DB   int
}

type MysqlConf struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}
