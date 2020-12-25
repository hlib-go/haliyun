package haliyunoss

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"testing"
)

func TestA(t *testing.T) {
	t.Log("111 ", config.Endpoint)

	bucket, _ := NewBucket(config)
	url, err := bucket.SignURL("images/20201225/ud.png", oss.HTTPGet, 3600)
	if err != nil {
		t.Error(err)
	}
	t.Log(url)
}

// Put 上传文件
func Put(bucket *oss.Bucket, key string, file []byte) (err error) {
	//bucket.Client.SetBucketACL()

	return bucket.PutObject(key, bytes.NewReader(file), oss.Option(nil))
}

// Get 下载文件
func Get(bucket *oss.Bucket, key string) {
	/*rc, err := bucket.GetObject(key)
	ioutil.ReadAll(rc)*/
}

// OssPutHandler 文件上传
func OssPutHandler(bucketFunc func() (*oss.Bucket, error)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
		var (
			err error
		)
		if request.Method != "POST" {
			err = errors.New("请使用POST请求上传文件")
			return
		}
		file, fh, err := request.FormFile("file")
		if err != nil {
			return
		}
		bucket, err := bucketFunc()
		if err != nil {
			return
		}
		objectKey := fh.Filename
		err = bucket.PutObject(objectKey, file)
		if err != nil {
			return
		}
	})
}

// OssGetHandler 文件下载
func OssGetHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

	})
}

// Rand32 使用crypto/rand 随机赋值byte数组， 然后md5返回32位十六进制字符串
func Rand32() string {
	var b = make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}
