package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("etc/order.ini") // 执行 `go run` 的相对路径
	err := viper.ReadInConfig()          // 读取配置信息
	if err != nil {                      // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.AllKeys())

	viper.SetDefault("app.page_size", 100) // 设置默认值
	fmt.Println(viper.Get("app.page_size"))

	fmt.Println(viper.GetString("database.user"))
	viper.Set("database.user", "admin") // 配置覆盖
	fmt.Println(viper.GetString("database.user"))
}
