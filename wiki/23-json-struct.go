package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	serverIP   string // 注意json不会解析首字母小写的成员
}
type Servers struct {
	Servers []Server
}

func main() {

	s := Servers{[]Server{{"localhost", "127.0.0.1"}, {"beijin", "127.0.0.2"}}}
	// struct 转为 []byte，再转为 json string
	ss, _ := json.Marshal(s)
	fmt.Println(string(ss))

	var servers Servers
	str := `{"servers":[
		{"serverName":"Guangzhou_Base","serverIP":"127.0.0.1"},
		{"serverName":"Beijing_Base","serverIP":"127.0.0.2"}
	]}`

	// json 转 struct，注意先将字符串转化为 []byte
	err := json.Unmarshal([]byte(str), &servers)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(servers)

	for _, v := range servers.Servers {
		fmt.Printf("server_name: %s, server_ip: %s \n", v.ServerName, v.serverIP)
	}
}
