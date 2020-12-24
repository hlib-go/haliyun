package haliyunoss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func New(config *Config, options ...oss.ClientOption) (*oss.Client, error) {
	return oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret, options...)
}

func NewBucket(config *Config, bucket string) (*oss.Bucket, error) {
	client, err := New(config)
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucket)
}
