package dal

import (
	"github.com/benxinm/tiktok/cmd/follow/dal/cache"
	"github.com/benxinm/tiktok/cmd/follow/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
