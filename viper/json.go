package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	//viper.SetConfigFile("./etc/order.yaml") // 指定配置文件
	viper.SetConfigName("order")  // 配置文件名
	viper.AddConfigPath("./etc/") // 指定查找配置文件的路径
	viper.SetConfigType("json")   // 默认yaml
	err := viper.ReadInConfig()   // 读取配置信息
	if err != nil {               // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("server name:" + viper.GetString("name"))

	// 读取子配置文件，注意自配置文件类型必须与主配置一致
}
