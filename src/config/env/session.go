package env

import (
	"github.com/go-redis/redis"
	"time"
)

var Session = &struct {
	TokenKey       string
	DataKey        string
	Expiration     int64
	UpdateDuration time.Duration
	GetClientFunc  func() *redis.Client
}{
	TokenKey:       "sess_token",
	DataKey:        "gin-gonic/gin/sessdata",
	UpdateDuration: 30 * time.Minute, // 30minutes
	Expiration:     6 * 30 * 86400,   // 6months
}
