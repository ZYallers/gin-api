package abst

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	ege *gin.Engine
	ctx *gin.Context
}

func (c *Controller) SetEngine(ege *gin.Engine) *Controller {
	c.ege = ege
	return c
}

func (c *Controller) GetEngine() *gin.Engine {
	return c.ege
}

func (c *Controller) SetContext(ctx *gin.Context) *Controller {
	c.ctx = ctx
	return c
}

func (c *Controller) GetContext() *gin.Context {
	return c.ctx
}

// args 三个参数：
// 第一个是code int，代表状态码
// 第二个是msg string，代表信息
// 第三个是data gin.H，代表数据
func (c *Controller) Json(args ...interface{}) {
	switch len(args) {
	case 1:
		c.ctx.JSON(http.StatusOK, gin.H{"code": args[0], "msg": "Ok", "data": nil})
	case 2:
		c.ctx.JSON(http.StatusOK, gin.H{"code": args[0], "msg": args[1], "data": nil})
	case 3:
		c.ctx.JSON(http.StatusOK, gin.H{"code": args[0], "msg": args[1], "data": args[2]})
	default:
		c.ctx.JSON(http.StatusOK, args)
	}
}