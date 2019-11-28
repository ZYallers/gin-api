package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type GeneController struct {
	abs.Controller
}

const geneUri = "http://slim.hxsapp.com/other/Gene/"

func Gene() GeneController {
	c := GeneController{}
	c.Config = map[string]abs.MethodConfig{
		"GetYMRGeneAgressByCheckNum": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 根据ID获取基因检测数据
 */
func (c GeneController) GetGeneReportById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

/**
 * 获取当前登录用户的基因检测列表
 */
func (c GeneController) GetGeneReportList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

/**
 * 获取FTO及3+1基因报告列表(3.2.6)
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=12131984
 * @author ZYaller
 * @time 2018/4/16
 */
func (c GeneController) GetGenCheckList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

/**
 * 获取FTO及3+1基因报告分页列表(3.3.0)
 * @see http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=12137426
 * @author ZYaller
 * @time 2018/6/4
 */
func (c GeneController) GetGenCheckPageList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

/**
 * 保存基因检测知情同意书
 */
func (c GeneController) SaveGeneAgreement(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

/**
 * 获取基因检测知情同意书
 */
func (c GeneController) GetGeneAgreement(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

/**
 * 通过id获取基因知情同意书
 */
func (c GeneController) GetGeneAgreementById(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetGeneAgreementWiteReport(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetInformation(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) SetInformation(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetReportStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetReportByExpress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) SaveYMRGeneDetecJson(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) SaveYMRGeneAgreement(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetYMRReportStatus(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetYMRReportByCheckNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}

func (c GeneController) GetYMRGeneAgressByCheckNum(ctx *gin.Context) {
	c.ServiceRewrite(ctx, geneUri+tool.CurrentMethodName())
}
