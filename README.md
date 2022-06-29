# go 实验室

## 接入redis

> https://github.com/go-redis/redis

## 知识点
- [位运算](./wiki/01-位运算.go)
- [闭包](./wiki/cluse.go)
- [interface](./wiki/interface.go)
- [goto](./wiki/goto.go)
- [range 可以迭代 array, slice, map, chan](./wiki/range.go)

### 常用内置包
- [time](./wiki/pkg/time.go)
- [fmt.Scan 接收输入](./wiki/pkg/fmt.Scan接收输入信息.go)
- [os.open 按行读取文件](./wiki/pkg/os.Open-按行读取文件.go)
- [json、map 之间的转化](./wiki/pkg/json-map.go)
- [json、struct 之间的转化](./wiki/pkg/json-struct.go)
#### 接收命令参数的两种方式
- [flag](./wiki/pkg/flag-获取命令行参数.go)
- [os.Args](./wiki/pkg/os.Args-获取命令行参数.go)

## 试用功能
1. 扫描指定目录下的所有文件，保存文件名到文件中 `go run scan.go`
2. 固定扫描 `Download/xunlei` 目录下的所有文件 `go run scan_download.go`
3. 将历史扫描结果保存到文件 `have_save_file.txt` 中 `go run save_scan.go`




