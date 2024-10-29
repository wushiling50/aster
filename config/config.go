package config

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	_ "github.com/spf13/viper/remote"
)

var (
	Service       *service
	Snowflake     *snowflake
	Mysql         *mySQL
	Redis         *redis
	RabbitMQ      *rabbitMQ
	Elasticsearch *elasticsearch

	runtime_viper = viper.New()
)

func Init(path string, service string) {
	runtime_viper.SetConfigType("yaml")

	hlog.Infof("config path: %v\n", path)

	if err := runtime_viper.ReadRemoteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			hlog.Fatal("could not find config files")
		} else {
			hlog.Fatal("read config error: %v", err)
		}
		hlog.Fatal(err)
	}

	configMapping(service)

	hlog.Infof("all keys: %v\n", runtime_viper.AllKeys())

	// 持续监听配置
	runtime_viper.OnConfigChange(func(e fsnotify.Event) {
		hlog.Infof("config file changed: %v\n", e.String())
	})
	runtime_viper.WatchConfig()
}

func configMapping(srv string) {
	c := new(config)
	if err := runtime_viper.Unmarshal(&c); err != nil {
		hlog.Fatal(err)
	}

	Snowflake = &c.Snowflake
	Mysql = &c.MySQL
	Redis = &c.Redis
	RabbitMQ = &c.RabbitMQ
	Elasticsearch = &c.Elasticsearch
	Service = GetService(srv)
}

func GetService(srvname string) *service {
	addrlist := runtime_viper.GetStringSlice("services." + srvname + ".addr")

	return &service{
		Name:     runtime_viper.GetString("services." + srvname + ".name"),
		AddrList: addrlist,
		LB:       runtime_viper.GetBool("services." + srvname + ".load-balance"),
	}
}

// Init any essentials for ci testing
func InitForTest() {
	Snowflake = &snowflake{
		WorkerID:      0,
		DatancenterID: 0,
	}

	Mysql = &mySQL{
		Addr:     "127.0.0.1:3306",
		Database: "aster",
		Username: "aster",
		Password: "aster",
		Charset:  "utf8mb4",
	}

	RabbitMQ = &rabbitMQ{
		Addr:     "127.0.0.1:5672",
		Username: "aster",
		Password: "aster",
	}

	Redis = &redis{
		Addr:     "127.0.0.1:6379",
		Password: "aster",
	}
}
