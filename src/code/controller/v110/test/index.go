package v110

import (
	"code/app/abst"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type testIndex struct {
	abst.Controller
}

func NewTestIndex(e *gin.Engine) *testIndex {
	ti := &testIndex{}
	ti.SetEngine(e)
	return ti
}

func (ti *testIndex) CheckOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok", "data": gin.H{"is_debugging": gin.IsDebugging()}})
}

func (ti *testIndex) AllConfig(c *gin.Context) {
	time.Sleep(4 * time.Second)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "15", "data": gin.H{"url": c.Request.URL.String()}})
}

func (ti *testIndex) Main(c *gin.Context) {
	ti.GetEngine().LoadHTMLFiles("./view/http/test/index/main.html")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": c.Request.URL.String(),
	})
}
