package db

import (
	"github.com/benxinm/tiktok/config"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB
var SF *utils.Snowflake

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.GetMysqlDSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		panic(err)
	}
	if err = DB.Use(gormopentracing.New()); err != nil { //开启gorm链路追踪
		panic(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(constants.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)

	DB = DB.Table(constants.VideoTableName)
	if SF, err = utils.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID); err != nil {
		panic(err)
	}
}
