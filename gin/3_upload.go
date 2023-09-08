package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	path, _ := os.Getwd()
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		dst := path + "/" + file.Filename
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	/*
		curl -X POST http://localhost:8080/upload \
		  -F "upload[]=@/Users/appleboy/test1.zip" \
		  -F "upload[]=@/Users/appleboy/test2.zip" \
		  -H "Content-Type: multipart/form-data"
	*/
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/multiUpload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := path + "/" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	r.Run(":8080")
}
