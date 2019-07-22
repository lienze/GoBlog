package router

import (
	"html/template"
	"net/http"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "Hello World")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/login.html")
	t.Execute(w, "")
}

func InitRouter() error {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/login", loginPage)
	return nil
}
