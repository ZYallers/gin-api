package community

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type VoteController struct {
	abs.Controller
}

const voteUri = "http://community.hxsapp.com/user/vote/"

func Vote() VoteController {
	c := VoteController{}
	c.Config = map[string]abs.MethodConfig{
		"SaveUserVoteData": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 保持用户投票内容
 */
func (c VoteController) SaveUserVoteData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, voteUri+tool.CurrentMethodName())
}

/**
 * 查询用户投票结果
 */
func (c VoteController) SelectUserVoteResult(ctx *gin.Context) {
	c.ServiceRewrite(ctx, voteUri+tool.CurrentMethodName())
}
