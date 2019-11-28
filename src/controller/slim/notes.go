package slim

import (
	"github.com/gin-gonic/gin"
	"src/abs"
	"src/library/tool"
)

type NotesController struct {
	abs.Controller
}

const notesUri = "http://slim.hxsapp.com/notes/Diary/"

func Notes() NotesController {
	c := NotesController{}
	c.Config = map[string]abs.MethodConfig{
		"UserAdd":  {ControllerNameFirstUpper: true},
		"HomeList": {ControllerNameFirstUpper: true},
		"Delete": {ControllerNameFirstUpper: true},
		"Detail": {ControllerNameFirstUpper: true},
		"PlanList": {ControllerNameFirstUpper: true},
	}
	return c
}

func (c NotesController) UserAdd(ctx *gin.Context) {
	c.ServiceRewrite(ctx, notesUri+tool.CurrentMethodName())
}

func (c NotesController) HomeList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, notesUri+tool.CurrentMethodName())
}

func (c NotesController) Delete(ctx *gin.Context) {
	c.ServiceRewrite(ctx, notesUri+tool.CurrentMethodName())
}

func (c NotesController) Detail(ctx *gin.Context) {
	c.ServiceRewrite(ctx, notesUri+tool.CurrentMethodName())
}

func (c NotesController) PlanList(ctx *gin.Context) {
	c.ServiceRewrite(ctx, notesUri+tool.CurrentMethodName())
}
