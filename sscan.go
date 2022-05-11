package main

import (
	"fmt"
	"github.com/alexchen/test/Cache"
)

func main() {
	var path string
	fmt.Printf("请输入要查询的文件:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	match := "*" + path + "*"
	res := Cache.SScan("have_save_file",0, match, 0)
	fmt.Println(res)

	res2 := Cache.SScan("laravel_database_files",0, match, 0)
	fmt.Println(res2)
}