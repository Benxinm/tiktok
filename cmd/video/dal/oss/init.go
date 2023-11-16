package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/benxinm/tiktok/cmd/video/dal"
	"github.com/benxinm/tiktok/config"
)

func Init() {
	client, err := oss.New(config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret)
	if err != nil {
		panic("oss connection failed")
	}
	dal.OssClient = client
}
