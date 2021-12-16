package core

import (
	"github.com/ZYallers/zgin/libraries/mvcs"
	"github.com/go-redis/redis"
	"src/config/env"
)

type Redis struct {
	mvcs.Redis
}

var session mvcs.RdsCollector

func (r *Redis) GetSession() *redis.Client {
	return r.NewClient(&session, &env.Redis.Session)
}
