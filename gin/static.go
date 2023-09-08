package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// http://127.0.0.1:8081/more_static/dyc.jpeg
	router.StaticFS("/more_static", http.Dir("sources"))

	router.StaticFile("/favicon.ico", "./sources/dyc1.jpeg")

	router.GET("/download", downloadFile)

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8081")
}

func downloadFile(c *gin.Context) {
	// 获取文件路径
	filePath := "/Users/chentao/Downloads/wallpaper/0003.jpeg"
	// 设置文件名
	fileName := "1003.jpeg"
	// 设置文件头
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	// 下载文件
	c.File(filePath)
}
