package mysql_gorm

import (
	"gorm.io/gorm/logger"
	"strconv"
)

type Config struct {
	Path         string          // 服务器地址:端口
	Port         int             `json:",default=3306"` // 端口
	Config       string          // 高级配置
	Dbname       string          // 数据库名
	Username     string          // 数据库用户名
	Password     string          // 数据库密码
	MaxIdleConns int             // 空闲中的最大连接数
	MaxOpenConns int             // 打开到数据库的最大连接数
	LogMode      logger.LogLevel // 是否开启Gorm全局日志
	LogZap       bool            `json:"LogZap,optional,default=false"` // 是否通过zap写入日志文件
}

func (m *Config) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + strconv.Itoa(m.Port) + ")/" + m.Dbname + "?" + m.Config
}
