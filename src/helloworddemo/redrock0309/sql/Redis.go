package sql

import (
	"github.com/go-redis/redis"
)


type Cfg struct {
	Addrs    []string
	Pwd      string
	PoolSize int
	DB       int
}


func RedisConnect() *redis.Client {

	c := Cfg{}
	c.Addrs = append(c.Addrs, "123.17.165.145:6384")
	c.PoolSize = 10
	c.Pwd = ""
	c.DB = 1

	return redis.NewClient(&redis.Options{
		Addr:     c.Addrs[0],
		Password: c.Pwd,
		PoolSize: c.PoolSize,
		DB:       c.DB,
	})
}

