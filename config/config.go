package config

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	Server        *server
	Mysql         *mySql
	Snowflake     *snowflake
	Service       *service
	Etcd          *etcd
	OSS           *oss
	Redis         *redis
	runtime_viper = viper.New()
)

func InitLocal(path string, service string) {
	runtime_viper.SetConfigType("yaml")
	runtime_viper.AddConfigPath(path)
	klog.Infof("local config path: %v\n", path)
	if err := runtime_viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			klog.Fatal("could not find config files")
		} else {
			klog.Fatal("read config error")
		}
		klog.Fatal(err)
	}
	configMapping(service)
	runtime_viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("config file changed: %v\n", e.String())
	})
	runtime_viper.WatchConfig()
}

func configMapping(srv string) {
	c := new(config)
	if err := runtime_viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	Snowflake = &c.Snowflake
	Server = &c.Server
	Server.Secret = []byte(runtime_viper.GetString("server.jwt-secret"))
	Etcd = &c.Etcd
	Mysql = &c.MySql
	Redis = &c.Redis
	OSS = &c.OSS
	Service = GetService(srv)
}

func GetService(srvname string) *service {
	addrlist := runtime_viper.GetStringSlice("services." + srvname + ".addr")
	return &service{
		Name:     runtime_viper.GetString("services." + srvname + ".name"),
		AddrList: addrlist,
		LB:       runtime_viper.GetBool("service." + srvname + ".load-balance"),
	}
}

func GetMysqlDSN() string {
	if Mysql == nil {
		klog.Fatal("config not found")
	}
	dsn := strings.Join([]string{Mysql.Username, ":",
		Mysql.Password, "@tcp(", Mysql.Addr, ")/",
		Mysql.Database, "?charset=" + Mysql.Charset + "&parseTime=true"}, "")
	return dsn
}
