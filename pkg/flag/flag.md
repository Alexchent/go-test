# 接收外部输入的方式

## flag 包

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var conf string
	flag.StringVar(&conf, "c", "config.yaml", "配置文件")
	fmt.Println(conf)
}
```