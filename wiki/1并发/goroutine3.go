package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello, world"
	done <- true
}

func main() {
	go aGoroutine()
	<-done
	println(msg)
}
