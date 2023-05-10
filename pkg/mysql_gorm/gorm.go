package mysql_gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDefaultGormDb new a default gorm db,generally used in svc.NewServiceContext(c config.Config) *ServiceContext
func NewDefaultGormDb(config Config) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       config.Dsn(), // DSN data source name
		DefaultStringSize:         255,          // string 类型字段的默认长度
		SkipInitializeWithVersion: false,        // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogMode),
	})
	if err != nil {
		panic("mysql connect failed :" + mysqlConfig.DSN)
	} else {
		temp, _ := db.DB()
		temp.SetMaxIdleConns(config.MaxIdleConns)
		temp.SetMaxOpenConns(config.MaxOpenConns)
	}
	return db
}

func NewGormDbWithDriverConfig(driverConfig mysql.Config, config Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(driverConfig), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogMode),
	})
	if err != nil {
		panic("mysql connect failed :" + driverConfig.DSN)
	} else {
		temp, _ := db.DB()
		temp.SetMaxIdleConns(config.MaxIdleConns)
		temp.SetMaxOpenConns(config.MaxOpenConns)
	}
	return db
}
