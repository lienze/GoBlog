package router

import (
	"GoBlog/src/config"
	"GoBlog/src/zdata"
	"GoBlog/src/zversion"
	"html/template"
	"net/http"
	"strconv"
)

//var ContentShow []string = []string{}

func rootPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
	r.ParseForm()
	t.Execute(w, zdata.IndexPage)
}

func showpost(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/showpost.html")
	r.ParseForm()
	zdata.PageShow.WebTitle = zdata.AllPageData.WebTitle
	r.ParseForm()
	//fmt.Println("showpost:", r.Form["name"][0])
	filePath := "./post/" + r.Form["name"][0]
	postID := zdata.GetPostIDFromPath(filePath)
	indexInfo := zdata.AllPostData[postID]
	zdata.PageShow.PageTitle = indexInfo.PostTitle
	zdata.PageShow.PageDate = indexInfo.PostDate
	zdata.PageShow.PageContent = indexInfo.PostContent
	zdata.PageShow.PageReadNum = indexInfo.PostReadNum
	zdata.PageShow.PageComments = indexInfo.PostComments
	zdata.PageShow.BlogVersion = zversion.Ver
	// update PostCommentNum
	tmp := zdata.AllPostData[postID]
	tmp.PostCommentNum = len(tmp.PostComments)
	zdata.AllPostData[postID] = tmp
	zdata.PageShow.PageCommentNum = indexInfo.PostCommentNum
	//fmt.Println(postID)
	t.Execute(w, zdata.PageShow)
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
		} else if iCurPage >= zdata.AllPageData.MaxPage {
			iCurPage = zdata.AllPageData.MaxPage
		}
	}
	zdata.CurPageData.CurPage = iCurPage
	zdata.CurPageData.MaxPage = zdata.AllPageData.MaxPage
	zdata.CurPageData.WebTitle = zdata.AllPageData.WebTitle
	// insert current page
	zdata.CurPageData.ContentShow = make([]string, 0)
	allDataLen := len(zdata.AllPageData.ContentShow)
	iPerPage := config.GConfig.PageCfg.MaxItemPerPage
	for i := (iCurPage - 1) * iPerPage; i < iCurPage*iPerPage && i < allDataLen; i++ {
		zdata.CurPageData.ContentShow =
			append(zdata.CurPageData.ContentShow, zdata.AllPageData.ContentShow[i])
	}
	t.Execute(w, zdata.CurPageData)
}

// temporary solution
//func getShowDownJS(w http.ResponseWriter, r *http.Request) {
//	fileContent, _ := file.ReadFile("./html/showdown.min.js")
//	w.Write([]byte(fileContent))
//}

func InitRouter() error {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/content", contentPage)
	http.HandleFunc("/showpost", showpost)
	//http.HandleFunc("/showdown.min.js", getShowDownJS)

	// init static file service
	files := http.FileServer(http.Dir("./public/"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	return nil
}
