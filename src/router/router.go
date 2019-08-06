package router

import (
	"html/template"
	"net/http"
	"sort"
)

//var ContentShow []string = []string{}

type ContentStruct struct {
	ContentShow []string
	MaxPage     int
	CurPage     int
}

var ContentPage ContentStruct

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
	t.Execute(w, ContentPage)
}

// temporary solution
//func getShowDownJS(w http.ResponseWriter, r *http.Request) {
//	fileContent, _ := file.ReadFile("./html/showdown.min.js")
//	w.Write([]byte(fileContent))
//}

func InitRouter() error {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/content", contentPage)
	//http.HandleFunc("/showdown.min.js", getShowDownJS)

	// init static file service
	files := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	return nil
}

func RefreshContentShow(mapFiles map[string]string) {
	var mapkeys []string
	for k := range mapFiles {
		mapkeys = append(mapkeys, k)
	}
	//fmt.Println(mapkeys)
	sort.Sort(sort.Reverse(sort.StringSlice(mapkeys)))
	ContentPage.ContentShow = make([]string, 0)
	for _, val := range mapkeys {
		//fmt.Println(key, " ", val)
		ContentPage.ContentShow = append(ContentPage.ContentShow, mapFiles[val])
	}
	ContentPage.MaxPage = 5
	ContentPage.CurPage = 1
}
