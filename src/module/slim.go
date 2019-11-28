package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/slim"
	"src/library/tool"
)

type SlimModule struct {
	abs.Module
}

func Slim() SlimModule {
	return SlimModule{}
}

/**
 * TODO::Android Controller Upper, 解决方案：待APP强制升级版本时，让客户端替换！
 */
func (a SlimModule) slimManualUpper(rg *gin.RouterGroup, c *slim.ManualController) {
	gp := rg.Group("/Manual")
	{
		gp.POST("/saveManualTypes", c.SaveManualTypes)
		gp.POST("/getManualTypes", c.GetManualTypes)
		gp.POST("/getManualData", c.GetManualData)
		gp.POST("/saveManualData", c.SaveManualData)
		gp.POST("/saveManualDataToClock", c.SaveManualDataToClock)
		gp.POST("/getManualDataByType", c.GetManualDataByType)
		gp.POST("/setBloodData", c.GetBloodData)
		gp.POST("/saveBloodData", c.SaveBloodData)
	}
}

func (a SlimModule) slimGeneUpper(rg *gin.RouterGroup, c *slim.GeneController) {
	gp := rg.Group("/Gene")
	{
		gp.POST("/getYMRReportByCheckNum", c.GetYMRReportByCheckNum)
		gp.GET("/getYMRReportByCheckNum", c.GetYMRReportByCheckNum)
		gp.POST("/getYMRReportStatus", c.GetYMRReportStatus)
		gp.GET("/getYMRReportStatus", c.GetYMRReportStatus)
		gp.GET("/saveYMRGeneDetecJson", c.SaveYMRGeneDetecJson)
		gp.POST("/saveYMRGeneDetecJson", c.SaveYMRGeneDetecJson)
		gp.POST("/saveYMRGeneAgreement", c.SaveYMRGeneAgreement)
		gp.GET("/saveYMRGeneAgreement", c.SaveYMRGeneAgreement)
	}
}

/**
 * TODO::Android/IOS Controller Upper, 解决方案：待APP强制升级版本时，让客户端替换！
 */
func (a SlimModule) slimWeighingUpper(rg *gin.RouterGroup, c *slim.WeighingController) {
	gp := rg.Group("/Weighing")
	{
		gp.POST("/getInfoByType", c.GetInfoByType)
		gp.POST("/saveReport", c.SaveReport)
		gp.POST("/getReport", c.GetReport)
		gp.POST("/getReportByIdOrDate", c.GetReportByIdOrDate)
		gp.POST("/errorLog", c.ErrorLog)
		gp.POST("/checkAdc", c.CheckAdc)
	}
}

func (a SlimModule) slimBongUpper(rg *gin.RouterGroup, c *slim.BongController) {
	gp := rg.Group("/Bong")
	{
		gp.GET("/getUserSleepLastXDays", c.GetUserSleepLastXDays)
		gp.POST("/getUserSleepLastXDays", c.GetUserSleepLastXDays)
	}
}

func (a SlimModule) Group(eg *gin.Engine) {
	gp := eg.Group("/" + tool.CurrentFileName())
	{
		slimBong := slim.Bong()
		a.slimBongUpper(gp, &slimBong)

		slimManual := slim.Manual()
		a.slimManualUpper(gp, &slimManual)

		slimWeighing := slim.Weighing()
		a.slimWeighingUpper(gp, &slimWeighing)

		slimGene := slim.Gene()
		a.slimGeneUpper(gp, &slimGene)

		a.BindMethodOfController(gp,
			slim.Comment(),
			slim.Common(),
			slim.Evaluate(),
			slim.Exhibit(),
			slim.Food(),
			slim.Isp(),
			slim.Nbong(),
			slim.Plan(),
			/* 运维提供的API没有调用记录，暂时注释生活、食物、调查问卷、方案等会员模块
			slim.Life(),
			slim.LsFood(),
			slim.Questionnaire(),
			slim.Scheme(),
			slim.SchemeManage(),*/
			slim.Shake(),
			slim.SlimPage(),
			slim.Sport(),
			slim.Treadmill(),
			slim.Ways(),
			slim.Notes(),
			slimWeighing,
			slimManual,
			slimBong,
			slimGene,
			slim.Sleep(),
		)
	}
}
