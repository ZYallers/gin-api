package mall

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type CustomerController struct {
	abs.Controller
}

const customerUri = "http://mall.hxsapp.com/base/Customer/"

func Customer() CustomerController {
	c := CustomerController{}
	c.Config = map[string]abs.MethodConfig{
		"UpdateCustomerAddress":    {ControllerNameFirstUpper: true},
		"DeleteCustomerAddress":    {ControllerNameFirstUpper: true},
		"AddCustomerAddress":       {ControllerNameFirstUpper: true},
		"GetCustomerAddress":       {ControllerNameFirstUpper: true},
		"WebGetCustomerAddress":    {ControllerNameFirstUpper: true},
		"WebAddCustomerAddress":    {ControllerNameFirstUpper: true},
		"WebDeleteCustomerAddress": {ControllerNameFirstUpper: true},
		"WebUpdateCustomerAddress": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c CustomerController) UpdateCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) DeleteCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) AddCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) GetCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) WebGetCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) WebAddCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) WebDeleteCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}

func (c CustomerController) WebUpdateCustomerAddress(ctx *gin.Context) {
	c.ServiceRewrite(ctx, customerUri+tool.CurrentMethodName())
}
