package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	auth, ok := session.Values["authenticated"].(bool)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "口令：宝塔正河妖")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret) // curl -s --cookie "cookie-name=MTQ....." localhost:8080/secret
	http.HandleFunc("/login", login)   // curl -s -I localhost:8080/login
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}