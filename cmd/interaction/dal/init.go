package dal

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
