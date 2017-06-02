package models

import (
	"fmt"
	"github.com/h-tko/revel-base/libraries"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

var db *gorm.DB

// DB接続処理.
func InitDB() {
	dbconf, err := libraries.LoadConfig("database.conf")

	if err != nil {
		revel.ERROR.Fatalf("can not read database.conf file. %#v\n", err)
	}

	host, _ := dbconf.String("host")
	port, _ := dbconf.String("port")
	dbname, _ := dbconf.String("dbname")
	user, _ := dbconf.String("user")
	password, _ := dbconf.String("password")

	dbc := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	dbsource, err := gorm.Open("postgres", dbc)

	if err != nil {
		revel.ERROR.Fatalf("cannot connect database. %s", dbc)
	}

	db = dbsource
	db.LogMode(true)
}

// DB切断処理.
//
// deferかなんかでアプリケーションの最後に呼ばれるようにする
func CloseDB() {
	db.Close()
}
