package abst

import (
	"code/app/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"sync"
	"time"
)

type Model struct {
	TableName string
}

type mqlPSer struct {
	Pointer *gorm.DB
	Singler sync.Once
	Err     error
}

var (
	enjoythin mqlPSer
)

func openMysql(host, port, database, username, password string, debug bool) (*gorm.DB, error) {
	tcp := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database +
		"?charset=utf8&collation=utf8_general_ci&parseTime=true&loc=PRC&maxAllowedPacket=0"
	db, err := gorm.Open("mysql", tcp)
	if err != nil {
		return nil, err
	}
	if debug {
		db.LogMode(true)
	}
	return db, err
}

func (m *Model) GetEnjoythin() *gorm.DB {
	enjoythin.Singler.Do(func() {
		logger.Info("", "Open Enjoythin Mysql Connection")
		host := os.Getenv("mysql_host")
		port := os.Getenv("mysql_port")
		database := os.Getenv("mysql_database")
		username := os.Getenv("mysql_username")
		pwd := os.Getenv("mysql_password")
		if enjoythin.Pointer, enjoythin.Err = openMysql(host, port, database, username, pwd, false); enjoythin.Err != nil {
			panic(enjoythin.Err)
		} else {
			enjoythin.Pointer.DB().SetMaxOpenConns(8)
			enjoythin.Pointer.DB().SetMaxIdleConns(2)
			enjoythin.Pointer.DB().SetConnMaxLifetime(30 * time.Second)
		}
	})
	if enjoythin.Pointer == nil || enjoythin.Pointer.DB().Ping() != nil {
		logger.Info("", "Reopen Enjoythin Mysql Connection")
		enjoythin = mqlPSer{}
		return m.GetEnjoythin()
	}
	return enjoythin.Pointer
}
