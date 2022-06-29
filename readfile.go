package main

import (
	"bufio"
	"fmt"
	"github.com/alexchen/test/Cache/redis"
	scan "github.com/alexchen/test/ScanService"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("/Users/chentao/Downloads/down.txt")
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
			redis.SAdd(scan.CacheKey, data)
		}
	}
}
