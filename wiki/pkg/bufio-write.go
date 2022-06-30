package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wr("./test.txt", "123")
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
