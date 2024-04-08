// 读取环境变量的方式
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	a := os.Getenv("HOME")
	_ = os.Setenv("GIRL_FRIEND", "小龙女")
	b := os.Getenv("GIRL_FRIEND")
	fmt.Println(a, b)

	viper.AutomaticEnv()
	c := viper.Get("GOPROXY")
	fmt.Println(c)

}
