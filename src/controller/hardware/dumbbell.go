package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type DumbbellController struct {
	abs.Controller
}

const dumbbellUri = "http://hardware.hxsapp.com/device/Dumbbell/"

func Dumbbell() DumbbellController {
	c := DumbbellController{}
	return c
}

/**
 * 检测哑铃MAC是否被绑定
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) CheckBellIsBind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}

/**
 * 哑铃绑定
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) Bind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}

/**
 * 哑铃解绑
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) Unbind(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}

/**
 * 哑铃同步数据
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) SaveData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}

/**
 * 哑铃首页
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) Index(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}

/**
 * 哑铃数据曲线
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) GetSportListByXDays(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}

/**
 * 根据日期查看当天运动详情
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c DumbbellController) GetSportDetailByDate(ctx *gin.Context) {
	c.ServiceRewrite(ctx, dumbbellUri+tool.CurrentMethodName())
}
