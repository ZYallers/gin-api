package v110

import (
	"code/app/abst"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type index struct {
	abst.Controller
}

func Index(e *gin.Engine) *index {
	i := new(index)
	i.SetEngine(e)
	return i
}

func (i *index) CheckOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok", "data": gin.H{"is_debugging": gin.IsDebugging()}})
}

func (i *index) AllConfig(c *gin.Context) {
	time.Sleep(4 * time.Second)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "15", "data": gin.H{"url": c.Request.URL.String()}})
}

func (i *index) Main(c *gin.Context) {
	i.GetEngine().LoadHTMLFiles("./view/test/index/main.html")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": c.Request.URL.String(),
	})
}
