package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

var ossClient *oss.Client

//TODO 后期放在config里面获取

func Init() {
	client, err := oss.New("oss-cn-shanghai.aliyuncs.com", "LTAI5tEJMyW8bRA4jgSmZvPi", "7LJoIutd0RHnd8HutnHWwrtmkUubDZ")
	if err != nil {
		panic("oss connection failed")
	}
	ossClient = client
}
