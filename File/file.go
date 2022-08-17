package myFile

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func AppendContent(filename, content string) {
	//filename := fmt.Sprintf(SavePath, time.Now().Format("060102"))
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开文件失败处理
		log.Fatal(err.Error())
	} else {
		//buf := []byte(content)
		//fd.Write(buf)
		fd.WriteString(content + "\n")
	}

	defer fd.Close()
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
			//redis.SAdd(scan.CacheKey, data)
		}
	}
}
