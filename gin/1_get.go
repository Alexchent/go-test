package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var msg struct {
	Name    string `json:"user"`
	Message string
	Number  int
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()

	// 使用AsciiJSON 生成仅有 ASCII字符的JSON，非ASCII字符将会被转义
	r.GET("/AsciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/JSON", func(c *gin.Context) {
		// 输出 {"lang":"GO语言","tag":"\u003cbr\u003e"}
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/someJson2", func(c *gin.Context) {
		// 输出 {"lang":"GO语言","tag":"\u003cbr\u003e"}
		c.JSON(http.StatusOK, gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		})
	})

	r.GET("/someJson3", func(c *gin.Context) {
		msg2 := &msg
		msg2.Number = 100
		msg2.Name = "alex"
		msg2.Message = "hello"
		c.JSON(http.StatusOK, msg2)
	})

	// curl --location --request GET 'http://127.0.0.1:8080/bindQuery?name=alex&address=上海'
	r.GET("/bindQuery", func(c *gin.Context) {
		var person Person
		if c.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
		}
		//c.String(200, "Success")
		c.JSON(http.StatusOK, person)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
