package file

import (
	"bufio"
	"fmt"
	"github.com/alexchen/test/Cache"
	scan "github.com/alexchen/test/ScanService"
	"io"
	"os"
	"strings"
)

const SavePath = "have_save_file_%s.txt"

func AppendContent(filename, content string) {
	//filename := fmt.Sprintf(SavePath, time.Now().Format("060102"))
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

func readString(filename string) {
	//f, err := os.Open("/Users/chentao/Downloads/down.txt")
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil && io.EOF == err {
			break
		}
		data := strings.Trim(line, "\n")

		if data != "" {
			fmt.Println(data)
			// 写入导redis内
			Cache.SAdd(scan.CacheKey, data)
		}
	}
}
