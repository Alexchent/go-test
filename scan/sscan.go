package main

import (
	"fmt"
	"github.com/alexchen/go_test/cache/redis"
	"regexp"
	"strings"
)

func main() {
	var keyWord string
	fmt.Printf("请输入要查询的文件:\n")
	_, err := fmt.Scan(&keyWord)
	if err != nil {
		return
	}

	keyWord = strings.ToLower(keyWord)
	count := SearchFromRedisSet("have_save_file", keyWord)
	res := fmt.Sprintf("本次扫描发现 %d 个文件", count)
	fmt.Println(res)
}

func SearchFromRedisSet(key, path string) (count int) {
	res := redis.SMembers(key)
	//过滤掉特殊字符-和_
	reg, _ := regexp.Compile("-|_")
	path = reg.ReplaceAllString(path, "")
	for _, val := range res {
		a := reg.ReplaceAllString(val, "")
		if strings.Contains(strings.ToLower(a), strings.ToLower(path)) {
			fmt.Println(val)
			count++
		}
	}
	return
}
