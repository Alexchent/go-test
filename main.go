package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取 get 参数
		token := r.URL.Query().Get("token")
		log.Default()
		log.Println(token)
		fmt.Fprintf(w, "Welcome to my websit！")
	})

	// 静态资源
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "慕容秋荻",
			Lastname:  "小狄",
			Age:       17,
		}
		json.NewEncoder(w).Encode(peter)
	})

	http.ListenAndServe(":80", nil)
}
