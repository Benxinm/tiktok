package dal

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/benxinm/tiktok/cmd/video/dal/cache"
	"github.com/benxinm/tiktok/cmd/video/dal/db"
	myOss "github.com/benxinm/tiktok/cmd/video/dal/oss"
)

var OssClient *oss.Client

func Init() {
	db.Init()
	cache.Init()
	myOss.Init()
}
