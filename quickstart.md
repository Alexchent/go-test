## go quickstart

1. 创建一个项目文件夹你的go源码
```
cd 
mkdir demo
cd demo
```

2. 为你的代码开启依赖项跟踪
```
go mod init demo.com/demo
```

3. 在你的编辑器中写一个hello.go文件

你可以使用下面的代码
```golang
package main
import "fmt"

func main() {
    fmt.Println("hello world")
}
```
4. 运行一下你的代码 `go run hello.go`

##  调用外包中的代码
1. 找一个想要使用的外部包 rsc.io/quote
2. `import` 导入这个包，如下
```go
package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```
3. 加载这个包到mod中
```
$ go mod tidy 
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
```

4. 运行你的代码
```
go run test.go
```



