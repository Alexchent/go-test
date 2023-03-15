package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	//viper.SetConfigFile("./etc/order.yaml") // 指定配置文件
	viper.SetConfigName("order")  // 配置文件名
	viper.AddConfigPath("./etc/") // 指定查找配置文件的路径
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {             // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("server name:" + viper.GetString("name"))

	// 读取子配置文件
	secret := viper.GetString("SECRET_FILE")
	fmt.Println("secret_file:" + secret)
	if secret != "" {
		viper.SetConfigFile(secret) // 指定配置文件
		err := viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
	})

	db_ini := viper.GetStringMap("PORTAL_DB")
	fmt.Println(db_ini)
	db := viper.GetString("PORTAL_DB.write_db_database")
	fmt.Println(db)

	/*
		r := gin.Default()
		// 访问/version的返回值会随配置文件的变化而变化
		r.GET("/version", func(c *gin.Context) {
			//log.Println(viper.GetString("SECRET_FILE"))
			fmt.Println(viper.GetStringMap("CLICKHOUSE"))
			c.String(http.StatusOK, viper.GetString("version"))
		})

		if err := r.Run(
			fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
			panic(err)
		}

	*/
}
