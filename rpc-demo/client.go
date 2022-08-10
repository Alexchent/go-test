package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 传的参数
type Params struct {
	Width, Height int
}

type PersonParams struct {
	Name string
}

// 主函数
func main() {
	// 1.连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	// 2.调用方法
	// 面积
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)
	// 周长
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("周长：", ret)

	ret2 := ""
	err4 := conn.Call("Person.MyName", PersonParams{"小万"}, &ret2)
	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Println(ret2)
}
