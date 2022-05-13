package main

import (
	"fmt"
	"github.com/alexchen/test/Cache"
	"strings"
)

func main() {
	var path string
	fmt.Printf("请输入要查询的文件:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	res := Cache.SMembers("have_save_file")
	for _, val := range res {
		if strings.Contains(val, path) {
			fmt.Println(val)
		}
	}

	res2 := Cache.SMembers("laravel_database_files")
	for _, val := range res2 {
		if strings.Contains(val, path) {
			fmt.Println(val)
		}
	}
}
