package main

import "fmt"

func main() {
	a,b,c := echo()
	fmt.Println(a,b,c)

	fmt.Println(echo2())
}

func echo() (string, string, string) {
	return "golang", "php" , "java"
}

func echo2() (a int, b ,c string) {
	return 1, "php" , "java"
}

func echo3() string {
	return "hello"
}