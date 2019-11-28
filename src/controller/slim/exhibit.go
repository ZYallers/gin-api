package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type ExhibitController struct {
	abs.Controller
}

const exhibitUri = "http://slim.hxsapp.com/other/Exhibit/"

func Exhibit() ExhibitController {
	c := ExhibitController{}
	return c
}

/**
 * 展厅 - 保存身体报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:06:01+0800
 */
func (c ExhibitController) SaveReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, exhibitUri+tool.CurrentMethodName())
}

/**
 * 展厅 - 通过id获取身体报告
 * @Author   lifeng
 * @DateTime 2017-09-16T14:09:24+0800
 */
func (c ExhibitController) GetReportById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, exhibitUri+tool.CurrentMethodName())
}

/**
 * 展厅 - 发短信
 */
func (c ExhibitController) SendSMS(ctx *gin.Context) {
	c.ServiceRewrite(ctx, exhibitUri+tool.CurrentMethodName())
}
