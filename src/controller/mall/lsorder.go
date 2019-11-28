package mall

import (
    "github.com/gin-gonic/gin"
    "src/abs"
    "src/library/tool"
)

type LsorderController struct {
    abs.Controller
}

const (
    lsorderUri = "http://mall.hxsapp.com/lsmember/Lsorder/"
)

func Lsorder() LsorderController {
    c := LsorderController{}
    c.Config = map[string]abs.MethodConfig{
        "GetOrderList":            {ControllerNameFirstUpper: true},
        "GetOrderStatusCounter":   {ControllerNameFirstUpper: true},
        "GetOrderDetailList":      {ControllerNameFirstUpper: true},
        "GetLSOrderDetailList":    {ControllerNameFirstUpper: true},
        "Whetheropenelics":        {ControllerNameFirstUpper: true},
        "Applyelicsinvoice":       {ControllerNameFirstUpper: true},
        "Openelicsinvoice":        {ControllerNameFirstUpper: true},
        "OrderNum":                {ControllerNameFirstUpper: true},
        "Comment":                 {ControllerNameFirstUpper: true},
        "HideOrder":               {ControllerNameFirstUpper: true},
        "AfterService":            {ControllerNameFirstUpper: true},
        "GetAferService":          {ControllerNameFirstUpper: true},
        "OrderConsultant":         {ControllerNameFirstUpper: true},
        "InvoiceInfo":             {ControllerNameFirstUpper: true},
        "GetAdviserUserOrderList": {ControllerNameFirstUpper: true},
        "DelGoodsRedNodeState":    {ControllerNameFirstUpper: true},
        "GetAdviserGoodsInfo":     {ControllerNameFirstUpper: true},
        "DelAdviserGoodsInfo":     {ControllerNameFirstUpper: true},
    }

    return c
}

func (c LsorderController) DelAdviserGoodsInfo(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) GetAdviserGoodsInfo(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) DelGoodsRedNodeState(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) GetAdviserUserOrderList(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) InvoiceInfo(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) OrderConsultant(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) GetAferService(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) AfterService(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) HideOrder(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) Comment(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) OrderNum(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) Applyelicsinvoice(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}
func (c LsorderController) Openelicsinvoice(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}
func (c LsorderController) Whetheropenelics(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) GetLSOrderDetailList(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) GetOrderDetailList(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

func (c LsorderController) GetOrderList(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}

// 我的订单各状态数量
func (c LsorderController) GetOrderStatusCounter(ctx *gin.Context) {
    c.ServiceRewrite(ctx, lsorderUri+tool.CurrentMethodName())
}
