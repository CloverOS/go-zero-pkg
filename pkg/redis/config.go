package redis

type Config struct {
	Addr     string `json:"addr" yaml:"addr"`         //地址
	Password string `json:"password" yaml:"password"` //密码
	DB       int    `json:"db" yaml:"db"`             //数据库
}
