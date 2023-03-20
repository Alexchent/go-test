package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var cfg *ini.File

func main() {
	Cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("fail to read file: %v", err)
		os.Exit(1)
	}
	cfg = Cfg

	secs := cfg.SectionStrings()
	fmt.Println(secs)

	// 读取配置，默认分区使用空字符串表示
	app_mode := cfg.Section("").Key("app_mode").String()
	fmt.Println("app_mode: " + app_mode)
	data_path := cfg.Section("paths").Key("data").String()
	fmt.Println("path.data:" + data_path)

	// 候选值限制
	fmt.Println("Server.Protocol:", cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 如果读取的值不再候选列表内，则返回默认值
	fmt.Println("Email.Protocol:", cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 自动类型转换
	fmt.Println("server.http_port: ", cfg.Section("server").Key("http_port").MustInt(999))
	fmt.Println("server.enforce_domain: ", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 修改配置文件
	cfg.Section("").Key("app_mode").SetValue("production")
	err = cfg.SaveTo("my-local.ini")
	if err != nil {
		panic(err)
	}
}
