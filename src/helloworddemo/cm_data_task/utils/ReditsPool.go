package utils

import (
	"github.com/go-redis/redis"
)

var (
	RedisPoolClient *redis.Client
)

type Cfg struct {
	Addrs    []string
	Pwd      string
	PoolSize int
	DB       int
}

func init() {

	RedisPoolClient = Connect()
}

func Connect() *redis.Client {

	c := Cfg{}
	c.Addrs = append(c.Addrs, IniParserInstance.GetString("redis", "adds"))
	c.PoolSize = int(IniParserInstance.GetInt64("redis", "poolsize"))
	c.Pwd = IniParserInstance.GetString("redis", "pwd")
	c.DB = int(IniParserInstance.GetInt64("redis", "db"))

	return redis.NewClient(&redis.Options{
		Addr:     c.Addrs[0],
		Password: c.Pwd,
		PoolSize: c.PoolSize,
		DB:       c.DB,
	})
}
