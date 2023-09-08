package config

import (
	_ "embed"
	"fmt"
)

//go:embed version.txt
var version string

func Version() {
	fmt.Println("the version:" + version)
}
