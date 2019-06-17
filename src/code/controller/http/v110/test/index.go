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
	ti := new(testIndex)
	ti.Egn = e
	return ti
}

func (ti *testIndex) CheckOk(c *gin.Context) {
	ti.SetContext(c)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok", "data": gin.H{"debug": gin.IsDebugging()}})
}

func (ti *testIndex) AllConfig(c *gin.Context) {
	ti.SetContext(c)
	time.Sleep(4 * time.Second)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "15", "data": gin.H{"url": c.Request.URL.String()}})
}

func (ti *testIndex) Main(c *gin.Context) {
	ti.SetContext(c)
	ti.Egn.LoadHTMLFiles("./view/http/test/index/main.html")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": c.Request.URL.String(),
	})
}