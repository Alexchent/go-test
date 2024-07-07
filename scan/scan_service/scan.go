package scan

import (
	"fmt"
	myFile "github.com/alexchen/go_test/File"
	"github.com/alexchen/go_test/cache/redis"
	"io/ioutil"
	path "path/filepath"
	"time"
)

// Start 单线程处理
func Start(path string) {
	start := time.Now()
	defer fmt.Println("扫描完毕，共耗时：", time.Since(start))

	scanFile2(path)
}

func scanFile2(filePath string) {
	fileInfoList, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println(err)
	}

	for i := range fileInfoList {
		fileName := fileInfoList[i].Name()
		if fileInfoList[i].IsDir() {
			scanFile2(filePath + "/" + fileName)
		} else {
			// 过滤文件
			ext := path.Ext(fileName)
			if ext == ".js" || ext == ".torrent" || ext == ".DS_Store" {
				continue
			}

			fullFilename := filePath + "/" + fileName
			if redis.SAdd(CacheKey, fullFilename) == 1 {
				fmt.Println("发现新的文件：", fullFilename)
				myFile.AppendContent("have_save_file.txt", fullFilename)
			}
		}
	}
}
