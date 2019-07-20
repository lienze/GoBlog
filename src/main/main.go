package main

import (
	"fmt"
	"github.com/lienze/go2db/dao"
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

func main() {
	server := http.Server{
		Addr: "10.0.2.15:8080",
	}
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/login", loginPage)
	fmt.Println("GoBlog is running...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	dao.InitDB("mytest")
}
