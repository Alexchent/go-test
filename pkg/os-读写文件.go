package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	rd()
	//wrb("china")
}

func rd() {
	content, err := os.ReadFile("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

// wrb 覆盖写
func wrb(content string) {
	message := []byte(content)

	err := os.WriteFile("hello", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
