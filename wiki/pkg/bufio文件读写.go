package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//wr("./test.txt", "123")
	readLine()
}

func wr(file string, content string) {
	file_os, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file_os.Close()

	writer := bufio.NewWriter(file_os)
	writer.WriteString(content + "\n")
	writer.Flush()
}

func readLine() {
	f, err := os.Open("./test.txt")
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
		fmt.Println(data)
	}
}
