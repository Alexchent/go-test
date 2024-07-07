package myFile

import (
	"bufio"
	"fmt"
	"io"
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
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = fd.WriteString(content + "\n")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func AppendContent2(fd *os.File, content string) {
	//fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	_, err := fd.WriteString(content + "\n")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// CreateDateDir basePath是固定目录路径
func CreateDateDir(folderPath string) (dirPath string) {
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

// ScanDir 遍历目录下的所有文件
func ScanDir(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		filename := file.Name()
		fullFilename := dir + "/" + filename
		// 忽略影藏文件
		if filename[0:1] == "." {
			continue
		}
		ext := path.Ext(filename)
		// 忽略指定后缀的文件
		if strings.ToLower(ext) == ".png" {
			continue
		}
		if file.IsDir() {
			ScanDir(fullFilename)
		} else {
			fmt.Println(fullFilename)
		}
	}
}

// ScaleDir2 扫描目录2
func ScaleDir2(dir string) error {
	//获取当前目录下的所有文件或目录信息
	err := path.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
