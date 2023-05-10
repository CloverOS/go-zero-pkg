package mongo

type Config struct {
	Url      string //连接地址
	DataBase string //数据库名称
	Password string //密码
	Enable   bool   //是否启用
}
