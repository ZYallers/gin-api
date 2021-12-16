package helper

import "github.com/gin-gonic/gin"

// QueryPostForm
func QueryPostForm(ctx *gin.Context, keys ...string) string {
	if len(keys) == 0 {
		return ""
	}
	if val, ok := ctx.GetQuery(keys[0]); ok {
		return val
	}
	if val, ok := ctx.GetPostForm(keys[0]); ok {
		return val
	}
	if len(keys) == 2 {
		return keys[1]
	}
	return ""
}
