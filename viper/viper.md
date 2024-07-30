# Viper 读取配置

> [Go配置管理神器——Viper中文教程](https://www.cnblogs.com/niuben/p/13896330.html)
> 
> [参考代码]()
## 安装 
```
go get github.com/spf13/viper
```

## 读取配置文件
### 方式一
如下所示：分别尝试从三个路径读取 `config.yaml` 文件
```go
viper.SetConfigName("config") // name of config file (without extension)
viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
viper.AddConfigPath(".")               // optionally look for config in the working directory
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("fatal error config file: %w", err))
}
```

### 方式二
```go
viper.SetConfigFile("./etc/order.json") // 指定配置文件路径
err := viper.ReadInConfig()             // 读取配置信息
if err != nil {                         // 读取配置信息失败
panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

### 多个配置文件合并，注意文件类型必须一致
```go
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
    err := viper.MergeInConfig()   // 合并配置文件
    if err != nil {
        panic(err)
    }
}
```
