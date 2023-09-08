package main

import (
	"example.com/hello/Model"
	"example.com/hello/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequest 从 JSON 或 form 绑定
type LoginRequest struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginJson(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err == nil {
		person := Model.FindByUserName(json.User)
		if json.Password == person.Email {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			fmt.Println("param", json)
			fmt.Println("data", person)
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func LoginForm(c *gin.Context) {
	var form LoginRequest
	// 这个将通过 content-type 头去推断绑定器使用哪个依赖。
	if err := c.ShouldBind(&form); err == nil {
		person := Model.FindByUserName(form.User)
		if form.Password == person.Email {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			fmt.Println(form)
			fmt.Println(person)
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func GetPerson(c *gin.Context) {
	p := Model.Find(3)
	c.JSON(http.StatusOK, p)
}

func main() {
	router := gin.Default()
	config.Version()

	// 绑定 JSON 的示例 ({"user": "manu", "password": "123"})
	/*
		curl --location --request POST 'http://127.0.0.1:8080/loginJSON' \
		--header 'Content-Type: application/json' \
		--data-raw '{"User":"menu","Password":"123"}'
	*/
	router.POST("/loginJSON", LoginJson)

	// 一个 HTML 表单绑定的示例 (user=manu&password=123)
	router.POST("/loginForm", LoginForm)

	router.GET("/person", GetPerson)

	// 监听并服务于 0.0.0.0:8080
	router.Run(":8080")
}
