# go 实验室

## 接入redis

> https://github.com/go-redis/redis

### 一、扫描指定目录下的所有文件，保存文件名到文件中
`go run scan.go`

### 二、固定扫描 `Download/xunlei` 目录下的所有文件
`go run scan_download.go`

### 三、将历史扫描结果保存到文件 `have_save_file.txt` 中
`go run save_scan.go`

## 知识点
- [位运算](./wiki/01-位运算.go)
- [闭包](./wiki/cluse.go)
- [interface](./wiki/interface.go)
- [goto](./wiki/goto.go)

### 常用内置包
- [flag 接收命令行参数](./wiki/pkg/flag.go)
- [time](./wiki/pkg/time.go)
- [fmt.Scan 接收输入](./wiki/pkg/01-fmt.Scan接收输入信息.go)
- [os.open 按行读取文件](./wiki/pkg/os.Open-按行读取文件.go)




