package haliyunoss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Put 上传文件
func Put(bucket *oss.Bucket, key string, file []byte) (err error) {
	return bucket.PutObject(key, bytes.NewReader(file), oss.Option(nil))
}

// Get 下载文件
func Get(bucket *oss.Bucket, key string) {
	/*rc, err := bucket.GetObject(key)
	ioutil.ReadAll(rc)*/
}
