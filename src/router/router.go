package router

import (
	"GoBlog/src/config"
	"html/template"
	"net/http"
	"sort"
	"strconv"
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
	r.ParseForm()
	//fmt.Println("contentPage:", r.Form["page"])
	// the para page may null, check it before use
	if r.Form["page"] == nil {
		//fmt.Println("contentPage page nil")
		ContentPage.CurPage = 1
	} else {
		iCurPage, err := strconv.Atoi(r.Form["page"][0])
		if err == nil {
			ContentPage.CurPage = iCurPage
		}
	}
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
	iMaxPage := len(mapkeys) / config.GConfig.PageCfg.MaxItemPerPage
	if len(mapkeys)%config.GConfig.PageCfg.MaxItemPerPage != 0 {
		iMaxPage++
	}
	if iMaxPage <= 0 {
		iMaxPage = 1
	}
	ContentPage.MaxPage = iMaxPage
	ContentPage.CurPage = 1
}
