package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Conn *gorm.DB

func InitDb() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/buer_management?charset=utf8&parseTime=True&loc=Local")
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(true)
	if err != nil {
		log.Fatal(err)
	}
	Conn = db
}
