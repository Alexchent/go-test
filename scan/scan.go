package main

import (
	"fmt"
	"github.com/alexchen/go_test/scan/scan_service"
	"os"
)

func main() {
	var path string
	fmt.Printf("请输入要扫描的目录:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if path == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		path = dir + "/Downloads"
	}
	scan.Start(path)
}
