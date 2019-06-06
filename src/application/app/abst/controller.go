package abst

import "github.com/gin-gonic/gin"

type Controller struct {
	Egn *gin.Engine
	Ctx *gin.Context
}
