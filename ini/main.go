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

	return
	// 读取配置，默认分区使用空字符串表示
	app_mode := cfg.Section("").Key("app_mode").String()
	fmt.Println(app_mode)
	data_path := cfg.Section("paths").Key("data").String()
	fmt.Println(data_path)

	// 候选值限制
	fmt.Println("Server Protocol:", cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 如果读取的值不再候选列表内，则返回默认值
	fmt.Println("Email Protocol:", cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 自动类型转换
	fmt.Println("port number: ", cfg.Section("server").Key("http_port").MustInt(999))
	fmt.Println("enforce domain: ", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 修改配置文件
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my-local.ini")

}

func testSection() {
	sec, _ := cfg.GetSection("")
	fmt.Println(sec)
}
