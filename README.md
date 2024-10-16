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
- [time 实现延迟的三种方式](pkg/time.go)
- [strconv](pkg/strconv-字符串转换.go)
- [rand 生成随机数](pkg/rand.go)
- [fmt.Scan 接收输入](pkg/fmt.Scan接收输入信息.go)
- [os.open 按行读取文件](pkg/os.Open-按行读取文件.go)
- [json、map 之间的转化](pkg/json-map.go)
- [json、struct 之间的转化](pkg/json-struct.go)
- [%v ,%+v, %#v 输出结构体时有什么区别](pkg/fmt.Printf.go)
- [hystrix 实现限流熔断](test/requestlimit_test.go)

## 使用第三方包
- [配置文件操作 viper](https://github.com/spf13/viper) 
- [命令行参数 pflag](wiki/pkg/pflag.go)

#### 接收命令参数的两种方式
- [flag](pkg/flag-获取命令行参数.go)
- [os.Args](pkg/os.Args-获取命令行参数.go)

## 踩坑
- [变量赋值](./wiki/01-变量声明中的坑.go)
- [data race](./wiki/data-race.go)

## 1.18版本新特性

### 1. [generic 泛型](./wiki/func/generics.go)
函数的参数和返回值支持泛型
### 2. [go work 工作空间模式](./workspace.md)

## 实用功能
4. [抽奖demo](./lottery)

## 实验
### [websocket](./gin-ws/websocket.md)
### [net/rpc](./rpc-demo)
golang官方的`net/rpc`库，使用`encoding/gob`进行编解码，由于其他语言不支持gob编解码方式，所以使用net/rpc库实现的RPC方法**没办法进行跨语言调用**
### [protobuf 作为restful接口传输协议](./protobuf.md)

### [设计模式](./SOILD)

参考文章：
> [为什么你不应该接收race代码](https://xargin.com/why-you-should-reject-racy-code/)