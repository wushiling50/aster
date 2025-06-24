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

// mysql:
//   addr: 127.0.0.1:3306
//   database: aster
//   username: aster
//   password: aster
//   charset: utf8mb4
