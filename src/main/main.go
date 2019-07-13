package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, "Hello World")
}

func main() {
	server := http.Server{
		Addr:"10.0.2.15:8080",
	}
	http.HandleFunc("/",rootPage)
	fmt.Println("GoBlog is running...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

