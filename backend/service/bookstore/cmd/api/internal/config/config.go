package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
	Email      struct {
		From     string
		Nickname string
		Secret   string
		Host     string
		Port     int64
		IsSSL    bool
	}
}
