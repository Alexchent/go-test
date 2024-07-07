package scan

import (
	"fmt"
	myFile "github.com/alexchen/go_test/File"
	"github.com/alexchen/go_test/cache/redis"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

func Do(path string) {

	start := time.Now()
	defer fmt.Println(time.Since(start))

	var wg sync.WaitGroup

	// 通道中存储文件名
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

	// 生产者
	wg.Add(1)
	scanFile(path, c, &wg)

	wg.Wait() // 等待所有goroutine结束
	fmt.Println("finish")
}

func scanFile(dir string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("正在扫描：", dir)

	var nextWg *sync.WaitGroup
	for i := range fileInfoList {
		fileName := fileInfoList[i].Name()
		fullFilename := dir + "/" + fileName
		if fileInfoList[i].IsDir() {
			// 开启多个扫描线程，扫描速度将远高于写入速度， 主线程结束时，go线程还没有完成
			nextWg.Add(1)
			go scanFile(fullFilename, c, nextWg)
		} else {
			// 过滤js和torrent文件
			ext := filepath.Ext(fileName) // 获取文件后缀
			if ext == ".DS_Store" || ext == ".js" || ext == ".torrent" {
				continue
			}
			// 文件名传入channel内
			c <- fullFilename
		}
	}
	nextWg.Wait()
}
