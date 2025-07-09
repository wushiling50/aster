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

type Snowflake struct {
	WorkerId      int64
	DatancenterId int64
}

type KqPusherConf struct {
	Brokers []string
	Topic   string
}

type DeepSeekModel struct {
	Endpoint    string
	Model       string
	MaxTokens   int
	Temperature float64
	TopP        float64
}
