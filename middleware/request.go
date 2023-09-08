package middleware

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiCostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.RequestURI
		var code string
		if c.Writer.Status() == http.StatusNotFound {
			path, code = "/404", "404"
		}
		start := time.Now()
		c.Next()
		if len(code) == 0 {
			code = strconv.Itoa(c.Writer.Status())
		}
		costTime := time.Since(start).Milliseconds()
		group := "api"

		log.Println(path, code, costTime, group)
	}
}
