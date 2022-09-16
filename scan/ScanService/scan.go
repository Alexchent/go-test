package scan

import (
	"fmt"
	"github.com/alexchen/go_test/Cache/redis"
	myFile "github.com/alexchen/go_test/File"
	"io/ioutil"
	"log"
	path "path/filepath"
	"sync"
	"time"
)

//const SavePath = "/Users/chentao/have_save_file_%s.txt"

const SavePath = "have_save_file_%s.txt"
const CacheKey = "have_save_file"

func Do(path string) {

	start := time.Now()
	defer fmt.Println(time.Since(start))

	var wg sync.WaitGroup

	c := make(chan string)
	defer close(c)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			content := <-c //阻塞直到取到数据
			filename := fmt.Sprintf(SavePath, time.Now().Format("060102"))
			myFile.AppendContent(filename, content)
			redis.SAdd(CacheKey, content)
		}
	}()

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

func scanFile(filePath string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileInfoList, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("正在扫描：", filePath)
	for i := range fileInfoList {

		fileName := fileInfoList[i].Name()

		if fileInfoList[i].IsDir() {
			// 开启多个扫描线程，扫描速度将远高于写入速度， 主线程结束时，go线程还没有完成
			wg.Add(1)
			go scanFile(filePath+"/"+fileName, c, wg)
		} else {
			// 过滤Mac的.DS_Store文件
			if fileInfoList[i].Name() == ".DS_Store" {
				continue
			}
			// 过滤js和torrent文件
			baseName := path.Base(fileName) // 获取文件名
			ext := path.Ext(baseName)       // 获取文件后缀
			if ext == ".js" || ext == ".torrent" {
				continue
			}

			// 文件名传入channel内
			body := filePath + "/" + fileName
			if redis.SAdd(CacheKey, body) == 1 {
				fmt.Println("发现新的文件：", fileName)
				c <- body
			}
		}
	}

	// 本目录扫描完毕
}

// scanFile2 单线程扫描
func scanFile2(filePath string) {
	fileInfoList, err := ioutil.ReadDir(filePath)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}

	fmt.Println("正在扫描：", filePath)

	for i := range fileInfoList {
		fileName := fileInfoList[i].Name()
		if fileInfoList[i].IsDir() {
			scanFile2(filePath + "/" + fileName)
		} else {
			// 过滤Mac的.DS_Store文件
			if fileInfoList[i].Name() == ".DS_Store" {
				continue
			}
			// 过滤js和torrent文件
			baseName := path.Base(fileName)
			ext := path.Ext(baseName)
			if ext == ".js" || ext == ".torrent" {
				continue
			}

			content := filePath + "/" + fileName
			if redis.SAdd(CacheKey, content) == 1 {
				fmt.Println("发现新的文件：", fileName)
				//filename := fmt.Sprintf(SavePath, time.Now().Format("060102"))
				myFile.AppendContent("have_save_file.txt", content)
			}
		}
	}
	// 本目录扫描完毕
}
