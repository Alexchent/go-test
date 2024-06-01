package main

import (
	"flag"
	"log"
)

var name string

// 子命令的实现
// go run flag-2.go go --name=alex
// go run flag-2.go php --name=alex
func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "name", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
