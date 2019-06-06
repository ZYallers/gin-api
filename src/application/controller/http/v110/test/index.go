package v110

import (
	"application/app/abst"
	"github.com/gin-gonic/gin"
	"net/http"
)

type testIndex struct {
	abst.Controller
}

func NewTestIndex(engine *gin.Engine) *testIndex {
	controller := new(testIndex)
	controller.Egn = engine
	return controller
}

func (this *testIndex) AllConfig(c *gin.Context) {
	this.Ctx = c
	/*this.Egn.LoadHTMLFiles("./view/index.html")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website Index",
	})*/
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "", "data": gin.H{"url": c.Request.URL.String()}})
}
