package Model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "root:wifi2021@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//fmt.Println(err.Error())
		panic(err.Error())
	}
	Db = db
	fmt.Println("数据库链接建立成功")
	Db.Debug()
}
