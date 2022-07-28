package main

import (
	"net/http"
)

func wsHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hello"))

}

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("localhost:8080", nil)
}
