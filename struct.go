package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
    
    //Books{"go 语言", "www.runoob.com", "Go 语言教程", 6495407}
    a := Books{"GO yuyan", "www.run", "go test", 1}
    fmt.Println(a)
      //使用key=>value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})
    
      // 忽略的字段为 0 或 空
    fmt.Println(Books{title:"GO", subject:"hai"})


    fmt.Println(a.title)
}