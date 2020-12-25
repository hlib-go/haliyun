package haliyunoss

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// 用于测试的配置
var config = new(Config)

func init() {
	/*
		{
		  "endpoint" : "oss-cn-shanghai.aliyuncs.com",
		  "access_key_id" : "",
		  "access_key_secret" : ""
		}
	*/
	// 读取阿里云AccessKey，从阿里云后台获取
	akjson, err := ioutil.ReadFile("access-key.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(akjson, &config)
	if err != nil {
		log.Fatal(err.Error())
	}
}
