package router

import (
	"GoBlog/src/file"
	"html/template"
	"net/http"
)

var ContentShow []string = []string{}

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
	t.Execute(w, ContentShow)
}

// temporary solution
func getShowDownJS(w http.ResponseWriter, r *http.Request) {
	fileContent, _ := file.ReadFile("./html/showdown.min.js")
	w.Write([]byte(fileContent))
}

func InitRouter() error {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/content", contentPage)
	http.HandleFunc("/showdown.min.js", getShowDownJS)
	return nil
}
