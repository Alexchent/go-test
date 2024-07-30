package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	CheckContains()

	//Test()
}

func Test() {
	program, err := expr.Compile(`date("2006-01-02")`)
	//program, err := expr.Compile(`duration("1h").Seconds() == 3600`)
	if err != nil {
		panic(err)
	}
	output, err := expr.Run(program, nil)
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}

func CheckContains() {
	program, err := expr.Compile(`Title contains "Hello" || Title contains "happy"`, expr.AsBool())
	if err != nil {
		panic(err)
	}

	env := Book{
		Title:  "Hello cool",
		Author: "alex",
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}
