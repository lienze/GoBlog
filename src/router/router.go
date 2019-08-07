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

var CurPageData ContentStruct
var AllPageData ContentStruct

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
	var iCurPage int = 0
	var err error
	// the para page may null, check it before use
	if r.Form["page"] == nil {
		//fmt.Println("contentPage page nil")
		iCurPage = 1
	} else {
		iCurPage, err = strconv.Atoi(r.Form["page"][0])
		if err != nil {
			iCurPage = 1
		}
	}
	CurPageData.CurPage = iCurPage
	CurPageData.MaxPage = AllPageData.MaxPage
	// insert current page
	CurPageData.ContentShow = make([]string, 0)
	allDataLen := len(AllPageData.ContentShow)
	for i := (iCurPage - 1) * 3; i < iCurPage*3 && i < allDataLen; i++ {
		CurPageData.ContentShow = append(CurPageData.ContentShow, AllPageData.ContentShow[i])
	}
	t.Execute(w, CurPageData)
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
	AllPageData.ContentShow = make([]string, 0)
	for _, val := range mapkeys {
		//fmt.Println(key, " ", val)
		AllPageData.ContentShow = append(AllPageData.ContentShow, mapFiles[val])
	}
	iMaxPage := len(mapkeys) / config.GConfig.PageCfg.MaxItemPerPage
	if len(mapkeys)%config.GConfig.PageCfg.MaxItemPerPage != 0 {
		iMaxPage++
	}
	if iMaxPage <= 0 {
		iMaxPage = 1
	}
	AllPageData.MaxPage = iMaxPage
	AllPageData.CurPage = 1
}
