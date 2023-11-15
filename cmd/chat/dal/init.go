package dal

import (
	"github.com/benxinm/tiktok/cmd/chat/dal/cache"
	"github.com/benxinm/tiktok/cmd/chat/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
