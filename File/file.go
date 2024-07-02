package myFile

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	path "path/filepath"
	"strings"
)

func PrintFile(fullFilename string) {
	fmt.Println("fullFilename =", fullFilename)
	//获取文件名带后缀
	filenameWithSuffix := path.Base(fullFilename)
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	//获取文件后缀
	fileSuffix := path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix =", fileSuffix)

	//获取文件名
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)
}

func GetFileNameOnly(fullFilename string) string {
	//获取文件名带后缀
	filenameWithSuffix := path.Base(fullFilename)
	//获取文件后缀
	fileSuffix := path.Ext(filenameWithSuffix)
	//获取文件名
	return strings.TrimSuffix(filenameWithSuffix, fileSuffix)
}

func AppendContent(filename, content string) {
	//filename := fmt.Sprintf(SavePath, time.Now().Format("060102"))
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开文件失败处理
		log.Fatal(err.Error())
	} else {
		//buf := []byte(content)
		//fd.Write(buf)
		fd.WriteString(content + "\n")
	}
}

// CreateDateDir basePath是固定目录路径
func CreateDateDir(folderPath string) (dirPath string) {
	//folderName := time.Now().Format("2006-01-02")
	//folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.MkdirAll(folderPath, os.ModePerm)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

func readString(filename string) {
	//f, err := os.Open("/Users/chentao/Downloads/down.txt")
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil && io.EOF == err {
			break
		}
		data := strings.Trim(line, "\n")

		if data != "" {
			fmt.Println(data)
		}
	}
}
