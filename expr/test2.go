package main

import (
	"fmt"
	"github.com/antonmedv/expr"
)

func main() {
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
