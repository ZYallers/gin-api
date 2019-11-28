package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CommonController struct {
	abs.Controller
}

const commonUri = "http://community.hxsapp.com/common/common/"

func Common() CommonController {
	c := CommonController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveBatchRecommend": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * [shareCounts 分享计数]
 * @return [type] [description]
 */
func (c CommonController) ShareCounts(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}

/**
 * 引导页-批量关注用户,加入圈子
 */
func (c CommonController) SaveBatchRecommend(ctx *gin.Context) {
	c.ServiceRewrite(ctx, commonUri+tool.CurrentMethodName())
}
