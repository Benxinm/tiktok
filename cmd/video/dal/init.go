package dal

import (
	"github.com/benxinm/tiktok/cmd/user/dal/db"
	"github.com/benxinm/tiktok/cmd/video/dal/cache"
)

func Init() {
	db.Init()
	cache.Init()
}
