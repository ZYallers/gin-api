package bonus

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type WeightManagementController struct {
	abs.Controller
}

const weightManagementUri = "http://bonus.hxsapp.com/base/WeightManagement/"

func WeightManagementUri() WeightManagementController {
	c := WeightManagementController{}
	c.Config = map[string]abs.MethodConfig{}
	return c
}

func (c WeightManagementController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, weightManagementUri+tool.CurrentMethodName())
}
