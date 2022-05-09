package scan

import (
	"fmt"
	"github.com/alexchen/test/Cache"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const SavePath = "/Users/chentao/have_save_file_%s.txt"

//const SavePath = "have_save_file_%s.txt"
const CacheKey = "have_save_file"

func SaveCache() {
	members := Cache.SMembers(CacheKey)
	fmt.Println(members)

	for _, member := range members {
		AppendContent(member)
	}
}

func Do() {
	var path string
	fmt.Printf("请输入要扫描的目录:\n")

	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}

	start := time.Now()
	defer fmt.Println(time.Since(start))

	c := make(chan string)
	defer close(c)

	go func() {
		for {
			file := <-c
			AppendContent(file)
			Cache.SAdd(CacheKey, file)
		}
	}()

	scanFile(path, c)
}

func AppendContent(content string) {
	filename := fmt.Sprintf(SavePath, time.Now().Format("060102"))
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开文件失败处理
	} else {
		//buf := []byte(content)
		//fd.Write(buf)
		_, err := fd.WriteString(content)
		if err != nil {
			return
		}
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)
}

func scanFile(filePath string, c chan string) {

	fileInfoList, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			fmt.Println(fileInfoList[i].Name())

			// 开启多个扫描线程，扫描速度将远高于写入速度， 主线程结束时，go线程还没有完成
			scanFile(filePath+"/"+fileInfoList[i].Name(), c)
		} else {
			// 过滤Mac的.DS_Store文件
			if fileInfoList[i].Name() == ".DS_Store" {
				continue
			}
			// 过滤js和torrent文件
			//baseName := path.Base(fileInfoList[i].Name())
			//ext := path.Ext(baseName)
			//if ext != "js" && ext != "torrent" {
			//	continue
			//}
			body := filePath + "/" + fileInfoList[i].Name() + "\n"
			// 文件名传入channel内
			c <- body
		}
	}
}
