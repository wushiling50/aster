package db

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	gormopentracing "gorm.io/plugin/opentracing"

	"github.com/wushiling50/aster/pkg/constants"
	"github.com/wushiling50/aster/pkg/utils"
)

var (
	DB       *gorm.DB
	SF       *utils.Snowflake
	DBUser   *gorm.DB
	DBDomain *gorm.DB
)

func Init() {
	var err error

	dsn, err := utils.GetMysqlDSN()
	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true, // 禁用默认事务
			Logger:                 logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
		})

	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大连接时间

	DBUser = DB.Table(constants.UserTableName).WithContext(context.Background())
	DBDomain = DB.Table(constants.DomainTableName).WithContext(context.Background())

	if SF, err = utils.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID); err != nil {
		panic(err)
	}
}
