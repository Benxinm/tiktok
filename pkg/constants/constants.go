package constants

import "time"

const (
	UserServiceName = "user"
	//mysql table name
	UserTableName   = "user"
	VideoTableName  = "video"
	FollowTableName = "follow"
	//limit
	MaxConnections     = 1000
	MaxQPS             = 100
	MaxIdleConnections = 10
	ConnMaxLifetime    = 10 * time.Second
	//jwt
	JwtSecrete = "jwt_secrete"
	//RPC
	MuxConnection     = 1
	RPCTimeout        = 3 * time.Second
	ConnectionTimeout = time.Second
)
