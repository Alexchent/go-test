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
	_, _ = file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)

	// 按照key排序
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		w.Write(m[key])
	}
	// 刷新缓冲
	w.Flush()
}

// GetFuncName 获取运行中的函数名
func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// InSlice 判断slice是否存在某个元素
func InSlice(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// ReverseSlice 反转slice中的元素
func ReverseSlice(slice []string) []string {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]
	}
	return slice
}

func TopK(slice []int, k int) []int {
	if len(slice) <= k {
		return slice
	}
	// 1. 构建小顶堆
	heap := slice[0:k]
	for i := k/2 - 1; i >= 0; i-- {
		adjustHeap(heap, i, k)
	}
	// 2. 遍历剩余元素
	for i := k; i < len(slice); i++ {
		if slice[i] > heap[0] {
			heap[0] = slice[i]
			adjustHeap(heap, 0, k)
		}
	}
	return heap
}

// 调整堆
func adjustHeap(heap []int, i, length int) {
	temp := heap[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && heap[k+1] < heap[k] {
			k++
		}
		if heap[k] < temp {
			heap[i] = heap[k]
			i = k
		} else {
			break
		}
	}
	heap[i] = temp
}

// ReverseString 反转字符串
func ReverseString(str string) string {
	if len(str) <= 1 {
		return str
	}
	return ReverseString(str[1:]) + str[:1]
}
