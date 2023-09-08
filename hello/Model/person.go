package Model

import (
	"fmt"
)

type Person struct {
	UserID   int64 `gorm:"primaryKey"`
	Username string
	Sex      string `gorm:"default:男"`
	Email    string
}

// TableName 指定Person的表名为person
func (Person) TableName() string {
	return "person"
}

func First() {
	// 获取第一条记录
	var firstUser Person
	Db.First(&firstUser)
	fmt.Println(firstUser)
}

func Last() {
	// 获取最后一条记录
	var lastUser Person
	Db.Last(&lastUser)
	fmt.Println(lastUser)
}

func All() {
	// 获取所有记录
	var users []Person
	Db.Find(&users)
	fmt.Println(users)
}

func FindByUserName(username string) Person {
	// where 条件查询 find查询多条 first查询第一条
	//var users []Person
	//Model.Db.Where("username=?", username).Find(&users)// select * from person where username="?"
	//fmt.Println(users)

	var user = Person{}
	Db.Where("username=?", username).First(&user) // select * from person where username="?" limit 1
	//fmt.Println(user)
	return user
}

func Find(id int) Person {
	// 根据主键查找
	var user Person
	Db.First(&user, id)
	fmt.Println(user)
	return user
}

func NewUser(user Person) {
	// 插入一条
	//user = Person{Username: "Jinzhu", Sex: "男", Email: "jinzhu@163.com", UserID: 1}
	Db.Create(&user) // 通过数据的指针来创建
	fmt.Println(user)
}

func BatchInsert(users []Person) {
	//批量插入
	//var users = []Person{{Username: "tom"},{Username: "alice"}}
	Db.Create(&users)
	for _, u := range users {
		fmt.Println(u.UserID)
	}
}
