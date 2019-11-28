package mall

import (
    "github.com/gin-gonic/gin"
    "src/abs"
    "src/library/tool"
)

type InvoiceController struct {
    abs.Controller
}

const (
    invoiceUri = "http://mall.hxsapp.com/lsmember/invoice/"
)

func Invoice() InvoiceController {
    c := InvoiceController{}

    return c
}

func (c InvoiceController) Elicsdlurl(ctx *gin.Context) {
    c.ServiceRewrite(ctx, invoiceUri+tool.CurrentMethodName())
}

func (c InvoiceController) Openelics(ctx *gin.Context) {
    c.ServiceRewrite(ctx, invoiceUri+tool.CurrentMethodName())
}

func (c InvoiceController) Openpaper(ctx *gin.Context) {
    c.ServiceRewrite(ctx, invoiceUri+tool.CurrentMethodName())
}
