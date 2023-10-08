package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
)

func AddCommas(n int64) string {
	s := strconv.FormatInt(n, 10)
	result := ""
	for i, c := range s {
		if i != 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}

func BuildCsvFile(m map[int][]string, filename string) {
	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create(filename + ".csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)

	// 按照key排序
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		w.Write(m[key])
		// 刷新缓冲
		w.Flush()
	}
}

// GetFuncName 获取运行中的函数名
func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
