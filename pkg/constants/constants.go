package constants

import "time"

const (
	UserServiceName = "user"
	//mysql table name
	UserTableName = "user"

	//limit
	MaxConnections     = 1000
	MaxQPS             = 100
	MaxIdleConnections = 10
	ConnMaxLifetime    = 10 * time.Second
	//jwt
	JwtSecrete = "jwt_secrete"
)
