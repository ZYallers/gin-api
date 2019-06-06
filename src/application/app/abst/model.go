package abst

import (
	"application/app/constant"
	"application/app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Model struct {
	Ctx       *gin.Context
	Enjoythin *gorm.DB
	TableName string
	Err       error
}

func openMysql(host, port, database, username, password string, debug bool) (*gorm.DB, error) {
	tcp := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&collation=utf8_general_ci&parseTime=true&loc=PRC&maxAllowedPacket=0"
	db, err := gorm.Open("mysql", tcp)
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	if debug {
		db.LogMode(true)
	}
	return db, err
}

func (this *Model) InitEnjoythin() {
	if this.Enjoythin != nil && this.Enjoythin.DB().Ping() == nil {
		return
	}
	host := os.Getenv("mysql_host")
	port := os.Getenv("mysql_port")
	database := os.Getenv("mysql_database")
	username := os.Getenv("mysql_username")
	pwd := os.Getenv("mysql_password")
	if this.Enjoythin, this.Err = openMysql(host, port, database, username, pwd, constant.DbDebug); this.Err != nil {
		panic(this.Err)
	}
	middleware.Recycling(this.Ctx, this.Enjoythin)
}
