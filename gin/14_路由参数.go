package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name", func(c *gin.Context) {
		//name := c.Param("name")
		names := c.Params
		c.String(http.StatusOK, "hello %s", names.ByName("name"))
	})

	r.GET("/search", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://cn.bing.com/")
	})

	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		if name == "chentao" && password == "123" {
			c.JSON(200, gin.H{
				"status":   "posted",
				"password": password,
				"name":     name,
			})

		} else {
			//c.String(http.StatusUnauthorized, "账号不存在")
			c.Redirect(http.StatusFound, "/user/who")
		}
	})

	//路由重定向
	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.Static("/assets", "./assets")
	r.StaticFile("/IMG_6580.jpg", "./assets/IMG_6580.jpg")
	r.Run(":8080")
}
