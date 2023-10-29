package main

import (
	"github.com/alexchen/go_test/rpc-demo/model"
	"log"
	"net/http"
	"net/rpc"
)

// 主函数
func main() {
	// 1.注册服务
	//rect := new(model.Rect)
	rect := &model.Rect{}

	chinese := model.Person{Name: "马冬梅", Age: 20}
	// 注册一个rect的服务
	rpc.Register(rect)
	rpc.Register(&chinese)
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
