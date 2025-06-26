package utils

import (
	"strings"

	"github.com/wushiling50/aster/config"
)

// GetMysqlDSN 会拼接 mysql 的 DSN
func GetMysqlDSN(mysql config.MysqlConf) string {
	dsn := strings.Join([]string{
		mysql.Username, ":", mysql.Password,
		"@tcp(", mysql.Addr, ")/",
		mysql.Database, "?charset=" + mysql.Charset + "&parseTime=true",
	}, "")

	return dsn
}
