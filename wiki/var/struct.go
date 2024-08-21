package main

import (
	"fmt"
	"strings"
)

type Books struct {
	title   string
	book_id int
	gender  bool
}

func main() {
	var books *Books
	fmt.Println(books) // nil
	fmt.Println(books == nil)

	book := new(Books) // 等同于 book := &Books{}
	fmt.Println(book)  // &{0, false}
	fmt.Println(book == nil)

	book2 := &Books{}
	fmt.Println(book2) // &{0, false}
	fmt.Println(book2 == nil)

	fmt.Println(strings.Repeat("-", 10))
	var book1 Books
	fmt.Println(book1) // {0 false}

	book1.title = "Python 教程"
	book1.book_id = 6788
	printBook(book1)

	// make 只能用于初始化 slice 、 map 、chan

	//Books{"go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	a := Books{"GO yuyan", 1, true}
	fmt.Println(a)

	//使用key=>value 格式
	fmt.Println(Books{title: "Go 语言", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "GO"})

	fmt.Println(a.title)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
