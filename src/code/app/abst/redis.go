package abst

import (
	"code/app/logger"
	goredis "github.com/go-redis/redis"
	"os"
	"sync"
)

type Redis struct {
}

type rdsPser struct {
	Pointer *goredis.Client
	Singler sync.Once
	Err     error
}

var cache rdsPser

func newRedisClient(host, port, password string) (*goredis.Client, error) {
	rds := goredis.NewClient(&goredis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return rds, nil
}

func (r *Redis) GetCache() *goredis.Client {
	cache.Singler.Do(func() {
		logger.Info("", "New Cache Redis Client")
		host := os.Getenv("redis_host")
		port := os.Getenv("redis_port")
		pwd := os.Getenv("redis_password")
		if cache.Pointer, cache.Err = newRedisClient(host, port, pwd); cache.Err != nil {
			panic(cache.Err)
		} else if _, cache.Err = cache.Pointer.Ping().Result(); cache.Err != nil {
			panic(cache.Err)
		}
	})
	if cache.Pointer == nil && cache.Pointer.Ping().Err() != nil {
		logger.Info("", "Renew Cache Redis Client")
		cache = rdsPser{}
		return r.GetCache()
	}
	return cache.Pointer
}
