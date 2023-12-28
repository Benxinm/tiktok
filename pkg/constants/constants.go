package constants

import "time"

const (
	UserServiceName        = "user"
	VideoServiceName       = "video"
	InteractionServiceName = "interaction"
	FollowServiceName      = "follow"
	ChatServiceName        = "chat"
	//mysql table name
	UserTableName           = "user"
	VideoTableName          = "video"
	FollowTableName         = "follow"
	CommentTableName        = "comment"
	FavoriteTableName       = "favorite"
	VideoFavouriteTableName = "video_favourite"
	ChatTableName           = "chat"
	MessageTableName        = "message"
	//Redis
	FollowRedis      = 1
	InteractionRedis = 2
	VideoRedis       = 3
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
	// snowflake
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0
)
