package hardware

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ShoesController struct {
	abs.Controller
}

const shoesUri = "http://hardware.hxsapp.com/device/shoes/Main/"

func Shoes() ShoesController {
	c := ShoesController{}
	c.Config = map[string]abs.MethodConfig{
		"GetUserLastSyncTime": {ControllerNameFirstUpper: true},
		"SaveBodyExamData":    {ControllerNameFirstUpper: true},
		"GetBodyExamList":     {ControllerNameFirstUpper: true},
		"GetBodyExamById":     {ControllerNameFirstUpper: true},
		"SaveShoesSportData":  {ControllerNameFirstUpper: true},
		"GetShoesSportList":   {ControllerNameFirstUpper: true},
		"GetShoesRunLogById":  {ControllerNameFirstUpper: true},
		"GetUserTotalInfo":    {ControllerNameFirstUpper: true},
		"GetUpgradePackage":   {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 用户上次同步时间戳
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112563
 * @author chs
 */
func (c ShoesController) GetUserLastSyncTime(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 同步身体测试数据（实时模式）
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112567
 * @author chs
 */
func (c ShoesController) SaveBodyExamData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 身体测试报告列表
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112571
 * @author chs
 */
func (c ShoesController) GetBodyExamList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 某个身体测试报告（走路或跑步）
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112575
 * @author chs
 */
func (c ShoesController) GetBodyExamById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 同步跑步鞋数据（普通模式）
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112768
 * @author chs
 */
func (c ShoesController) SaveShoesSportData(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 跑步鞋记录列表
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112774
 * @author chs
 */
func (c ShoesController) GetShoesSportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 某个跑步结果详情
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112777
 * @author chs
 */
func (c ShoesController) GetShoesRunLogById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 跑步鞋累计历史
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112779
 * @author chs
 */
func (c ShoesController) GetUserTotalInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}

/**
 * 获取最新的固件升级安装包
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=17112583
 * @author chs
 */
func (c ShoesController) GetUpgradePackage(ctx *gin.Context) {
	c.ServiceRewrite(ctx, shoesUri+tool.CurrentMethodName())
}
