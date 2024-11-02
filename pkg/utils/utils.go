package utils

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/wushiling50/aster/config"
)

func TimeParse(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func GetMysqlDSN() (string, error) {
	if config.Mysql == nil {
		return "", errors.New("config not found")
	}

	dsn := strings.Join([]string{config.Mysql.Username, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true"}, "")

	return dsn, nil
}

func GetMQUrl() string {
	if config.RabbitMQ == nil {
		hlog.Fatal("config not found")
	}

	url := strings.Join([]string{"amqp://", config.RabbitMQ.Username, ":", config.RabbitMQ.Password, "@", config.RabbitMQ.Addr, "/"}, "")

	return url
}

func GetEsHost() (string, error) {
	if config.Elasticsearch == nil {
		return "", errors.New("elasticsearch not found")
	}

	return config.Elasticsearch.Host, nil
}

func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}

	l.Close()

	return true
}

func GetAvailablePort() (string, error) {
	if config.Service.AddrList == nil {
		return "", errors.New("utils.GetAvailablePort: config.Service.AddrList is nil")
	}
	for _, addr := range config.Service.AddrList {
		if ok := AddrCheck(addr); ok {
			return addr, nil
		}
	}
	return "", errors.New("utils.GetAvailablePort: not available port from config")
}
