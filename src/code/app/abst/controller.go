package abst

import "github.com/gin-gonic/gin"

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
