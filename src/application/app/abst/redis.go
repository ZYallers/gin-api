package abst

import (
	"application/app/middleware"
	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis"
	"os"
)

type Redis struct {
	Ctx   *gin.Context
	Cache *goredis.Client
	Err   error
}

func newRedisClient(host, port, password string) (*goredis.Client, error) {
	rds := goredis.NewClient(&goredis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	if _, err := rds.Ping().Result(); err != nil {
		return nil, err
	}
	return rds, nil
}

func (this *Redis) InitCache() {
	if this.Cache != nil && this.Cache.Ping().Err() == nil {
		return
	}
	host := os.Getenv("redis_host")
	port := os.Getenv("redis_port")
	pwd := os.Getenv("redis_password")
	if this.Cache, this.Err = newRedisClient(host, port, pwd); this.Err != nil {
		panic(this.Err)
	}
	middleware.Recycling(this.Ctx, this.Cache)
}
