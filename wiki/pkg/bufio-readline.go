package main

// 按行读取文件内容
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	readLine()
}

func readLine() {
	f, err := os.Open("/Users/chentao/Downloads/user_fee.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') // 返回str
		//line, _, err := rd.ReadLine()  // 返回 []byte
		if err != nil && io.EOF == err {
			break
		}
		data := strings.Trim(line, "\n")

		if data != "" {
			fmt.Println(data)
		}
	}
}

func readStr() {
	f, err := os.Open("/Users/chentao/Downloads/user_fee.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine() // 返回 []byte
		if err != nil && io.EOF == err {
			break
		}
		fmt.Println(string(line))
	}
}
