package upload

type Config struct {
	LocalDir string    `json:",optional"` //本地上传路径
	Oss      OssConfig `json:",optional"` //Oss上传配置
}

type OssConfig struct {
	Enable          bool   `json:",optional"`
	AccessKeyId     string `json:",optional"`
	AccessKeySecret string `json:",optional"`
	Endpoint        string `json:",optional"`
	BucketName      string `json:",optional"`
}
