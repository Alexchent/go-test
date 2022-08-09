package main

var apple string
var done bool

func setup() {
	apple = "hello, world"
	done = true
}

// go run -race data-race.go 检查程序中是否存在data race问题
func main() {
	go setup()
	for !done {
	}
	print(apple)
}
