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
	WebTitle    string
}

type IndexStruct struct {
	PostPath    string
	PostTitle   string
	PostProfile string
}

var CurPageData ContentStruct
var AllPageData ContentStruct
var IndexData IndexStruct

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
	var iCurPage int = 0
	var err error
	// the para page may null, check it before use
	if r.Form["page"] == nil {
		iCurPage = 1
	} else {
		iCurPage, err = strconv.Atoi(r.Form["page"][0])
		if err != nil {
			iCurPage = 1
		}
		if iCurPage <= 0 {
			iCurPage = 1
		} else if iCurPage >= AllPageData.MaxPage {
			iCurPage = AllPageData.MaxPage
		}
	}
	CurPageData.CurPage = iCurPage
	CurPageData.MaxPage = AllPageData.MaxPage
	CurPageData.WebTitle = AllPageData.WebTitle
	// insert current page
	CurPageData.ContentShow = make([]string, 0)
	allDataLen := len(AllPageData.ContentShow)
	iPerPage := config.GConfig.PageCfg.MaxItemPerPage
	for i := (iCurPage - 1) * iPerPage; i < iCurPage*iPerPage && i < allDataLen; i++ {
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
	files := http.FileServer(http.Dir("./public/"))
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
	AllPageData.WebTitle = config.GConfig.WebSite.WebTitle
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

func RefreshIndexData() {

}
