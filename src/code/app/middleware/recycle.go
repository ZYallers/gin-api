package middleware

import (
	"code/app/constant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"sync"
)

func Recycle(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var recycleMap sync.Map
		if val, ok := c.Get(constant.RecycleKey); ok {
			recycleMap = val.(sync.Map)
			recycleMap.Range(func(k, v interface{}) bool {
				switch v.(type) {
				case *gorm.DB:
					_ = v.(*gorm.DB).Close()
					if gin.IsDebugging() {
						logger.Info("*gorm.DB is closed.", zap.Any("v", v))
					}
				case *redis.Client:
					_ = v.(*redis.Client).Close()
					if gin.IsDebugging() {
						logger.Info("*redis.Client is closed.", zap.Any("v", v))
					}
				}
				return true
			})
		}
	}
}

func Recycling(c *gin.Context, value interface{}) {
	key := fmt.Sprintf("%p", &value)
	var recycleMap sync.Map
	if val, ok := c.Get(constant.RecycleKey); ok {
		recycleMap = val.(sync.Map)
		if _, ok := recycleMap.Load(key); !ok {
			recycleMap.Store(key, value)
		}
	} else {
		recycleMap.Store(key, value)
		c.Set(constant.RecycleKey, recycleMap)
	}
}
