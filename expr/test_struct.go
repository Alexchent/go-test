package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

// User 注意首字母必须是大写
type User struct {
	Channel    string
	Balance    int64
	Group      string
	ReadTime   int64
	ListenTime int64
	Level      int64
}

func main() {
	T1()
	T2()
	T3()
	T4()
}

func T1() {
	code := `Channel=="123" && Balance==9 || Group contains "shanghai"`
	program, _ := expr.Compile(code, expr.AsBool())
	out, err := expr.Run(program, User{
		Channel: "123",
		Balance: 9,
		Group:   "China:shanghai",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

// T2 加法
func T2() {
	code := `ReadTime+ListenTime > 100`
	program, _ := expr.Compile(code, expr.AsBool())
	out, err := expr.Run(program, User{
		ReadTime:   90,
		ListenTime: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

// T3 乘法
func T3() {
	code := `ReadTime*Level > 100`
	program, _ := expr.Compile(code, expr.AsBool())
	out, err := expr.Run(program, User{
		ReadTime: 50,
		Level:    3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

// T4 字符串串联
func T4() {
	code := `Channel+":"+Group`
	program, _ := expr.Compile(code)
	out, err := expr.Run(program, User{
		Channel: "huaw",
		Group:   "China:shanghai",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
