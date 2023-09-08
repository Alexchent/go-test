package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	/*
		curl --location --request POST 'http://127.0.0.1:8080/post?id=1&page=10' \
		--form 'name="alex"' \
		--form 'message="hello"'
	*/
	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"name":    name,
		})
	})

	/**
	curl --location --request POST 'http://127.0.0.1:8080/form_post' \
	--form 'message="hello"'
	*/
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run(":8080")
}
