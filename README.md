# go 实验室
> https://golang.google.cn/doc/tutorial/getting-started

## 接入redis

> https://github.com/go-redis/redis

## 知识点
- [位运算](./wiki/01-位运算.go)
- [闭包](./wiki/func/cluse.go)
- [interface](./wiki/interface.go)
- [goto](./wiki/goto.go)
- [range 可以迭代 array, slice, map, chan](./wiki/range.go)

### 常用内置包
- [time 实现延迟的三种方式](./wiki/pkg/time.go)
- [strconv](./wiki/pkg/strconv-字符串转换.go)
- [rand 生成随机数](./wiki/pkg/rand.go)
- [fmt.Scan 接收输入](./wiki/pkg/fmt.Scan接收输入信息.go)
- [os.open 按行读取文件](./wiki/pkg/os.Open-按行读取文件.go)
- [json、map 之间的转化](./wiki/pkg/json-map.go)
- [json、struct 之间的转化](./wiki/pkg/json-struct.go)
- [%v ,%+v, %#v 输出结构体时有什么区别](./wiki/pkg/fmt.Printf.go)
- [hystrix 实现限流熔断](test/requestlimit_test.go)

## 使用第三方包
- [配置文件操作 viper](https://github.com/spf13/viper) 

#### 接收命令参数的两种方式
- [flag](./wiki/pkg/flag-获取命令行参数.go)
- [os.Args](./wiki/pkg/os.Args-获取命令行参数.go)

## 高级
### 保障所有 goroutine 完整执行的技巧
- [sync.Group]()

## 踩坑
- [变量赋值](./wiki/01-变量声明中的坑.go)
- [data race](./wiki/data-race.go)

## 1.18版本新特性

### 1. [generic 泛型](./wiki/func/generics.go)
函数的参数和返回值支持泛型
### 2. [go work 工作空间模式](./workspace.md)

## 实用功能
1. 扫描指定目录下的所有文件，保存文件名到文件中 `go run scan.go`
2. 固定扫描 `Download/xunlei` 目录下的所有文件 `go run scan_download.go`
3. 将历史扫描结果保存到文件 `have_save_file.txt` 中 `go run save_scan.go`


## 实验
### [websocket](./gin-ws/websocket.md)
### [net/rpc](./rpc-demo)
golang官方的`net/rpc`库，使用`encoding/gob`进行编解码，由于其他语言不支持gob编解码方式，所以使用net/rpc库实现的RPC方法**没办法进行跨语言调用**
### [protobuf 作为restful接口传输协议](./protobuf.md)

参考文章：
> [为什么你不应该接收race代码](https://xargin.com/why-you-should-reject-racy-code/)