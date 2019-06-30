package db

import (
	"fmt"
	"log"
	"net/url"

	"github.com/app/realworld/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

/*
DBとのconnectionを返す。
初期化がすでに行われていれば、初期化を行わずにgormを返す。
*/
func DBConn() *gorm.DB {
	// dbのconfigの呼び出し。
	conf := config.DBConfig

	if db != nil {
		return db
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName)

	// NOTE: https://github.com/bxcodec/go-clean-arch/blob/master/main.go#L42
	val := url.Values{}
	val.Add("parseTime", "1")
	// NOTE: 一旦Localに設定。
	val.Add("loc", "Local")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	_db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db = _db

	return db
}
