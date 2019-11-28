package module

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/controller/mall"
	"src/library/tool"
)

type MallModule struct {
	abs.Module
}

func Mall() MallModule {
	return MallModule{}
}

func (a MallModule) Group(eg *gin.Engine) {

	gp := eg.Group("/" + tool.CurrentFileName())
	{
		a.BindMethodOfController(gp,
			mall.Bind(),
			mall.Cart(),
			mall.Category(),
			mall.Coupon(),
			mall.Customer(),
			mall.DailyRec(),
			mall.Goods(),
			mall.GroupBuy(),
			mall.GymOrder(),
			mall.GymUniformPay(),
			mall.Invoice(),
			mall.Logistic(),
			mall.Lsorder(),
			mall.Lsuser(),
			mall.Market(),
			mall.Market400(),
			mall.NewGoods(),
			mall.Order(),
			mall.Prepayment(),
			mall.Reward(),
			mall.Spike(),
			mall.UniformPay(),
			mall.UserAccount(),
			mall.Video(),
			mall.Alipay(),
			mall.Wxpay(),
		)
	}
}
