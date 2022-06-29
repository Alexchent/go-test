package main

import (
	"fmt"
	"github.com/alexchen/test/Cache/redis"
	"strings"
)

func main() {
	var path string
	fmt.Printf("请输入要查询的文件:\n")
	_, err := fmt.Scan(&path)
	if err != nil {
		return
	}
	path = strings.ToLower(path)
	SearchFromRedisSet("have_save_file", path)

	SearchFromRedisSet("laravel_database_files", path)

	//res := Cache.SMembers("have_save_file")
	//for _, val := range res {
	//	// 处理分割符
	//	val = strings.ToLower(val)
	//	if strings.Contains(val, path) {
	//		fmt.Println(val)
	//	}
	//}

	//res2 := Cache.SMembers("laravel_database_files")
	//for _, val := range res2 {
	//	if strings.Contains(strings.ToLower(val), path) {
	//		fmt.Println(val)
	//	}
	//}
}

func SearchFromRedisSet(key, path string) {
	res := redis.SMembers(key)
	for _, val := range res {
		if strings.Contains(strings.ToLower(val), path) {
			fmt.Println(val)
		}
	}
}
