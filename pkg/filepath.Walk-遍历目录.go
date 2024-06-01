package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// go run filepath.Walk-遍历目录.go /Users/chentao/Documents/转正
func main() {
	// os.Args 获取命令行参数
	list := os.Args
	if len(list) < 2 {
		log.Fatal("请输入扫描文件路径")
	}
	path := os.Args[1]
	fmt.Println(path)

	fd, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil {
		log.Println(err.Error())
	}

	// 遍历 path
	filepath.Walk(path,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				log.Println("dir:", path)
				return nil
			}
			ss := strings.Split(path, "/")

			// 过滤 .DS_Store 文件
			if ss[len(ss)-1] == ".DS_Store" {
				return nil
			}

			log.Println("file:", path)

			buf := []byte(path + "\n")
			fd.Write(buf)
			return nil
		})
}
