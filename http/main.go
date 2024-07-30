package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// http://localhost/?token=123
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取 get 参数
		token := r.URL.Query().Get("token")
		log.Println(token)
		fmt.Fprintf(w, "Welcome to my websit！")
	})

	// 静态资源
	// http://localhost/static
	//fs := http.FileServer(http.Dir("static/"))
	fs := http.FileServer(http.Dir("/Users/chentao/Documents/wallpaper"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 解析json数据
	// curl --location 'http://localhost/decode' --header 'Content-Type: application/json' --data '{"age":10,"firstname":"zhao","lastname":"liner"}'
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age)
	})

	// 返回json
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "慕容秋荻",
			Lastname:  "小狄",
			Age:       17,
		}
		json.NewEncoder(w).Encode(peter)
	})
	// 将路由和处理函数绑定起来
	http.HandleFunc("/hello", HandleHelloRequest)

	http.Handle("/hi", &HelloHandler{content: "Hi World!"})

	fmt.Println("web 启动, 端口号：80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err.Error())
	}
}

func HandleHelloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// HelloHandler 实现了http.Handler接口的结构体
type HelloHandler struct {
	content string
}

// 实现了ServeHTTP方法的结构体，就是一个处理器
func (r *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "content=%s, args=%v", r.content, req.URL)
}
