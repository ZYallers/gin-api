package env

import (
	"github.com/ZYallers/zgin/app"
	"os"
)

var Redis = &struct {
	Session app.RedisClient
}{}

func init() {
	switch App.Env {
	case DevMode, GrayMode:
		Redis.Session = app.RedisClient{Host: os.Getenv("redis_session_host"),
			Port: os.Getenv("redis_session_port"), Pwd: os.Getenv("redis_session_password"), Db: 0}
	case ProdMode:
		Redis.Session = app.RedisClient{Host: "xxxxxx", Port: "6379", Pwd: "xxxxxx", Db: 0}
	}
}
