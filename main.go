package main

import (
	"fmt"
	myFile "github.com/alexchen/go_test/File"
)

func main() {
	var path string
	fmt.Printf("请输入要查询的文件:\n")
	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	myFile.ScanDir(path)
}
