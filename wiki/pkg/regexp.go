package main

import (
	"fmt"
	"regexp"
)

func main() {
	//str := "ssis-616"
	compile, _ := regexp.Compile("-|_")

	str := compile.ReplaceAllString("xxxxssis_616", "")
	fmt.Println(str)
}
