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

    var book1 Books
    book1.title = "Python 教程"
    book1.author = "jack"
    book1.subject = "test"
    book1.book_id = 6788

    printBook(book1)
}



func printBook(book Books) {
	fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}