package haliyunoss

// 阿里云配置，从阿里云平台获取
type Config struct {
	Endpoint        string `json:"endpoint"`          // OSS endpoint
	AccessKeyID     string `json:"access_key_id"`     // AccessId
	AccessKeySecret string `json:"access_key_secret"` // AccessKey
}
