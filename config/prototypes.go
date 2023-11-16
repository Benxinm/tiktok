package config

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type snowflake struct {
	WorkerID     int64
	DatacenterID int64
}

type service struct {
	Name     string
	AddrList []string
	LB       bool
}

type mySql struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type etcd struct {
	Addr string
}

type rabbitmq struct {
	Addr     string
	Username string
	Password string
}

type redis struct {
	Addr     string
	Password string
}

type oss struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	ManDirection    string
}

type config struct {
	Server    server
	Snowflake snowflake
	MySql     mySql
	Etcd      etcd
	Rabbitmq  rabbitmq
	Redis     redis
	OSS       oss
}
