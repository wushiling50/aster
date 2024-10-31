package constants

import "time"

const (
	// service name
	APIServiceName = "api"

	// for db
	UserTableName         = "user"
	DomainTableName       = "domain"
	MaxIdleConns          = 10
	MaxConnections        = 1000
	ConnMaxLifetime       = 10 * time.Second
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// for cache
	RedisDBScore  = 1
	RedisDBNation = 2
)
