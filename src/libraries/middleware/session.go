package middleware

import (
	"github.com/ZYallers/zgin/libraries/tool"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"src/config/env"
	"src/libraries/helper"
	"time"
)

const keyPrefix = "ci_session:"

func UpdateSessionExpireHandler(ctx *gin.Context) {
	ctx.Next()
	go func() {
		defer tool.SafeDefer()

		token := helper.QueryPostForm(ctx, env.Session.TokenKey)
		vars := sessionData(token)
		if vars == nil {
			return
		}

		//logger.Use("session_update_expire").Info(keyPrefix+token, zap.String("flag", "Ready"), zap.Any("vars", vars))

		client := getSessionClient()
		if client == nil {
			return
		}

		var lastRegenerate int
		if val, ok := vars["__ci_last_regenerate"].(int); ok {
			lastRegenerate = val
		} else {
			return
		}

		now := time.Now()
		if now.After(time.Unix(int64(lastRegenerate), 0).Add(env.Session.UpdateDuration)) {
			vars["__ci_last_regenerate"] = now.Unix()
			if ciVars, ok := vars["__ci_vars"].(map[string]interface{}); ok {
				newCiVars := map[string]interface{}{}
				for k := range ciVars {
					newCiVars[k] = now.Unix() + env.Session.Expiration
				}
				vars["__ci_vars"] = newCiVars
			}
			client.Set(keyPrefix+token, tool.PhpSerialize(vars), time.Duration(env.Session.Expiration)*time.Second)
			//logger.Use("session_update_expire").Info(keyPrefix+token, zap.String("flag", "Updated"), zap.Any("vars", vars))
		}
	}()
}

func getSessionClient() *redis.Client {
	if env.Session.GetClientFunc != nil {
		return env.Session.GetClientFunc()
	}
	return nil
}

func sessionData(token string) map[string]interface{} {
	if token == "" {
		return nil
	}

	client := getSessionClient()
	if client == nil {
		return nil
	}

	if str, _ := client.Get(keyPrefix + token).Result(); str != "" {
		return tool.PhpUnserialize(str)
	}

	return nil
}
