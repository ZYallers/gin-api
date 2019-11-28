package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type FoodController struct {
	abs.Controller
}

const foodUri = "http://slim.hxsapp.com/slim/food/"

func Food() FoodController {
	c := FoodController{}
	c.Config = map[string]abs.MethodConfig{
		"SearchFood": {ControllerNameFirstUpper: true},
	}
	return c
}

/**
 * 获取全部食物分类
 */
func (c FoodController) GetAllCategory(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 *  [saveRecord 保存饮食记录]
 */
func (c FoodController) SaveRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [search 搜索食物名称]
 */
func (c FoodController) Search(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 *  [deleteRecord 删除一条记录]
 */
func (c FoodController) DeleteRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [getNutrition 获取指定食物的营养成分]
 */
func (c FoodController) GetNutrition(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [getUnit 获取指定食物的规格]
 */
func (c FoodController) GetUnit(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [getTodayRecords 取当天饮食记录]
 */
func (c FoodController) GetTodayRecord(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [getRecords 取用户指定日期的饮食记录]
 */
func (c FoodController) GetRecords(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [saveCollection 收藏食品]
 */
func (c FoodController) SaveCollection(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [deleteCollection 取消收藏食品]
 */
func (c FoodController) DeleteCollection(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [deleteCustom 删除自定义食品]
 */
func (c FoodController) DeleteCustom(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [saveCustom 新增自定义食品]
 */
func (c FoodController) SaveCustom(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

/**
 * [getFoodInfo 获取指定食品详情]
 */
func (c FoodController) GetFoodInfo(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

func (c FoodController) GetUnitListByFoodId(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}

func (c FoodController) SearchFood(ctx *gin.Context) {
	c.ServiceRewrite(ctx, foodUri+tool.CurrentMethodName())
}
