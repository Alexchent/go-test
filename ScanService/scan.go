package scan

import (
	"fmt"
	"github.com/alexchen/test/Cache/redis"
	"io/ioutil"
	"log"
	"os"
	path "path/filepath"
	"sync"
	"time"
)

//const SavePath = "/Users/chentao/have_save_file_%s.txt"

const SavePath = "have_save_file_%s.txt"
const CacheKey = "have_save_file"

func SaveCache() {
	members := redis.SMembers(CacheKey)
	fmt.Println(members)

	for _, member := range members {
		AppendContent(member)
	}
}

func Do(path string) {

	start := time.Now()
	defer fmt.Println(time.Since(start))

	var wg sync.WaitGroup

	c := make(chan string)
	defer close(c)

	go func() {
		for {
			file := <-c //阻塞直到取到数据
			AppendContent(file)
			redis.SAdd(CacheKey, file)
		}
	}()

	wg.Add(1)
	scanFile(path, c, &wg)

	// 注意点 主程必须在go线程后结束
	wg.Wait() // 等待所有goroutine结束
	fmt.Println("finish")
}

// Start 单线程处理
func Start(path string) {
	start := time.Now()
	defer fmt.Println(time.Since(start))

	scanFile2(path)
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

func scanFile(filePath string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileInfoList, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("正在扫描：", filePath)
	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			//fmt.Println("正在扫描：", fileInfoList[i].Name())
			// 开启多个扫描线程，扫描速度将远高于写入速度， 主线程结束时，go线程还没有完成
			wg.Add(1)
			go scanFile(filePath+"/"+fileInfoList[i].Name(), c, wg)
		} else {
			// 过滤Mac的.DS_Store文件
			fmt.Println("发现文件：：", fileInfoList[i].Name())
			if fileInfoList[i].Name() == ".DS_Store" {
				continue
			}
			// 过滤js和torrent文件
			baseName := path.Base(fileInfoList[i].Name())
			ext := path.Ext(baseName)
			if ext == ".js" || ext == ".torrent" {
				continue
			}
			body := filePath + "/" + fileInfoList[i].Name() + "\n"
			// 文件名传入channel内
			c <- body
		}
	}

	// 本目录扫描完毕
}

func scanFile2(filePath string) {
	fileInfoList, err := ioutil.ReadDir(filePath)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println("正在扫描：", filePath)
	for i := range fileInfoList {
		if fileInfoList[i].IsDir() {
			scanFile2(filePath + "/" + fileInfoList[i].Name())
		} else {
			// 过滤Mac的.DS_Store文件
			if fileInfoList[i].Name() == ".DS_Store" {
				continue
			}
			// 过滤js和torrent文件
			baseName := path.Base(fileInfoList[i].Name())
			ext := path.Ext(baseName)
			if ext == ".js" || ext == ".torrent" {
				continue
			}

			body := filePath + "/" + fileInfoList[i].Name()
			res := redis.SAdd(CacheKey, body)
			if res == 1 {
				fmt.Println("发现新的文件：", fileInfoList[i].Name())
				AppendContent(body + "\n")
			}
		}
	}
	// 本目录扫描完毕
}
