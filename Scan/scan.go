package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var path string
	fmt.Printf("请输入要扫描的目录:\n")
	fmt.Scan(&path)

	start := time.Now()
	defer fmt.Println(time.Since(start))

	c := make(chan string)
	defer close(c)

	go func() {
		for {
			file := <-c
			appendContent(file)
		}
	}()

	saveFile(path, c)
}

func appendContent(content string) {
	fd, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开文件失败处理
	} else {
		buf := []byte(content)
		fd.Write(buf)
	}
	defer fd.Close()
}

func saveFile(path string, c chan string) {

	fileInfoList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			fmt.Println(fileInfoList[i].Name())

			// 开启多个扫描线程，扫描速度将远高于写入速度， 主线程结束时，go线程还没有完成
			saveFile(path+"/"+fileInfoList[i].Name(), c)
		} else {
			body := path + "/" + fileInfoList[i].Name() + "\n"
			// 文件名传入channel内
			c <- body
		}
	}
}
