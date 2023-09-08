package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	UserID   int64  `json:"user_id,primary"`
	Username string `json:"username,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Email    string `json:"email,omitempty"`
}

// TableName 指定表名为person    默认情况下表名是单数 people
func (Person) TableName() string {
	return "person"
}

func main() {
	// 链接数据库blog
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                          // default size for string fields
		DisableDatetimePrecision:  true,                                                                         // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                         // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                         // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                        // auto configure based on currently MySQL version
	}), &gorm.Config{})
	//db, err := gorm.Open("mysql", "root:123456@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	//db.SingularTable(true)
	//defer db.Close()
	//添加一条记录
	//user := Person{Username: "chentao", Sex: "男", Email: "doit@163.com"}
	//db.Create(&user)
	//db.Select("Username", "Email").Create(&user)

	var users = []Person{{Username: "jinzhu1"}, {Username: "jinzhu2"}, {Username: "jinzhu3"}}
	db.Create(&users)
	for _, user := range users {
		fmt.Println(user.UserID)
	}

	//查询第一条，按住键排序 SELECT * FROM users ORDER BY id LIMIT 1;
	//firstUser := new(Person)
	//var firstUser Person
	//db.First(&firstUser)
	//fmt.Println(firstUser)

	// 更新
	//db.Model(&firstUser).Update("email","jinzhu@163.com")

	//查询第一条不排序
	//var firstUser2 Person
	//db.Take(&firstUser2)
	//fmt.Println(firstUser2)

	//查询最后一条，按住键排序
	//var lastUser Person
	//db.Last(&lastUser)
	//fmt.Println(lastUser)

	// 获取所有的记录
	//var users[] Person
	//db.Find(&users)
	//fmt.Println(users)
}
