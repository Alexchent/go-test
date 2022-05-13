package main

import (
	"fmt"
	"github.com/alexchen/test/Cache"
	scan "github.com/alexchen/test/ScanService"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))

	data := Cache.SMembers("have_save_file")
	for _, v := range data {
		scan.AppendContent(strings.TrimRight(v, "\n") + "\n")
	}

	//data = Cache.SMembers("laravel_database_files")
	//for _,v := range data {
	//	scan.AppendContent(v + "\n")
	//}
}
