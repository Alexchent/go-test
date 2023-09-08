package main

import (
	"example.com/hello/Model"
	"fmt"
)

// 创建一个go.mod
// go mod init demo.com/greetings
// To do that, use the go mod edit command to edit the example.com/greetings module
// to redirect Go tools from its module path (where the module isn't) to the local directory (where it is
// go mod edit -replace demo.com/greetings=../greetings

func ini() {
	//
	//Person.First()
	//
	//Person.Last()
	//
	user := Model.Person{Username: "alex", Sex: "男", Email: "alex@163.com", UserID: 2}
	Model.NewUser(user)
	fmt.Println(user)
	//
	//var users = []Person.Person{{Username: "tom"},{Username: "alice"}}
	//Person.BatchInsert(users)

	//Person.Find(12)

	//person := Person.FindByUserName("alex")
	//fmt.Println(person)
}
