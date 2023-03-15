package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Log struct {
	Encoding string
}

type Config struct {
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
	Log     Log
}

var Conf = new(Config)

func main() {
	viper.SetConfigFile("./etc/order.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()             // 读取配置信息
	if err != nil {                         // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	fmt.Println(Conf)
}
