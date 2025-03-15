package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(dsn string) {
	//dsn := "root:root@tcp(192.168.211.151:3306)/Bolog_User?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic("fail to connect to database")
		log.Fatalln("fail to connect to database")
	}
	err = db.AutoMigrate(&UserBasic{})
	if err != nil {
		log.Fatalln("[DB ERROR]", err)
	}
	DB = db
}
