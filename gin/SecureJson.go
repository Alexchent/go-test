package main

import "net/http"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 你也可以使用自己的 SecureJSON 前缀
	r.SecureJsonPrefix("hello")
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"] 默认
		// 将输出：hello["lena","austin","foo"]  设置了前缀 hello
		c.SecureJSON(http.StatusOK, names)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
