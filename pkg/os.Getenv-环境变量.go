package main

import (
	"fmt"
	"os"
)

// linux/MacOS 设置仅限当前会话有效的环境变量 export TUTUS=123
func main() {
	fmt.Println(os.Getenv("HOME"))
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Getenv("GOROOT"))
	fmt.Println(os.Getenv("TUTUS"))
}
