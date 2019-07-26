package router

import (
	"html/template"
	"net/http"
)

var DaysOfWeek []string = []string{}

func rootPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "Hello World")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/login.html")
	t.Execute(w, "")
}

func contentPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/content.html")
	t.Execute(w, DaysOfWeek)
}

func InitRouter() error {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/content", contentPage)
	return nil
}
