package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

// User 注意首字母必须是大写
type User struct {
	Channel int64
	Balance int64
}

type UserInfo struct {
	Channel string
	Balance string
}

func main() {
	code := `Channel=="123" && Balance=="9"`

	program, _ := expr.Compile(code, expr.AsBool())
	out, err := expr.Run(program, User{
		Channel: 123,
		Balance: 9,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	out, err = expr.Run(program, UserInfo{
		Channel: "123",
		Balance: "9",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
