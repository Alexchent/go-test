package main

import (
	"fmt"
	"strconv"
)

//strconv包,字符转换：
//1) Append添加字符串元素:AppendInt,AppednBool,AppendQuote, AppendQuoteRune
//2) Format格式化程字符串:FormatBool(false),FormatInt(1234, 10),FormatUint(12345, 10),Itoa(1023)
//3) Parse字符串转换为其他类型:ParseBool("false")，ParseFloat("123.23", 64)，ParseInt("1234", 10, 64)，ParseUint("12345", 10, 64)，Atoi("1023")
//    返回值value, err := strconv.ParseBool("false")

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	a := strconv.AppendInt([]byte{}, 12, 10) //以10进制方式追加
	a = strconv.AppendQuote(a, "abc")        // 双引号
	a = strconv.AppendBool(a, false)
	a = strconv.AppendQuoteRune(a, '单') // 单引号
	//a = strconv.AppendFloat(a, "3.14")
	//fmt.Println(a)  // 输出slice
	fmt.Println(string(a)) // 转化为字符串

	// Format格式化字符串
	fmt.Println("=====================其他数据类型格式化为字符串=====================")
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(-123, 10)) // 转换为10进制数字符串
	fmt.Println(strconv.FormatFloat(3.14, 'f', -1, 64))
	fmt.Println(strconv.FormatUint(100, 10)) // uint 无符号整形

	fmt.Println(strconv.Itoa(666)) // FormatInt的简写

	// 字符串转换为其他类型
	fmt.Println("=================字符串转化为其他类型=================")
	fmt.Println(strconv.ParseFloat("3.14", 64))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseInt("125", 16, 64)) //由16进制转换为10进制 16^2+16^1*2+16^0*5 = 256 + 32 + 5
	fmt.Println(strconv.Atoi("123"))             // 等同于 ParseInt("123", 10, 64)
}
