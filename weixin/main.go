package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"weixin-demo/util"
)

// 与填写的服务器配置中的Token一致
const Token = "59284aa85709ddaf3bd246030060f6a2"

func main() {
	router := gin.Default()

	router.GET("/wx", WXCheckSignature)

	log.Fatalln(router.Run(":8088"))
}

// WXCheckSignature 微信接入校验
func WXCheckSignature(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := util.CheckSignature(signature, timestamp, nonce, Token)
	if !ok {
		log.Println("微信公众号接入校验失败!")
		return
	}

	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString(echostr)
}
