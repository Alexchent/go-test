package util

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

const (
	Unknown = "Unknown"
	UTF8    = "UTF-8"
	GBK     = "GBK"
)

// DetectEncoding 判断编码放肆
func DetectEncoding(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if utf8.Valid(data) {
		return UTF8
	}

	decoder := mahonia.NewDecoder("GBK")
	utf := decoder.ConvertString(string(data))
	if utf != "" {
		return GBK
	}

	return Unknown
}
