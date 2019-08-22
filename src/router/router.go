package router

import (
	"GoBlog/src/config"
	"GoBlog/src/file"
	"GoBlog/src/zdata"
	"GoBlog/src/ztime"
	"GoBlog/src/zversion"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")
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
		} else if iCurPage >= zdata.IndexPage.MaxPage {
			iCurPage = zdata.IndexPage.MaxPage
		}
	}
	zdata.SetCurIndexPageShow(iCurPage)
	t.Execute(w, zdata.IndexPage)
}

func showpost(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/showpost.html")
	r.ParseForm()
	zdata.PageShow.WebTitle = config.GConfig.WebSite.WebTitle
	//fmt.Println("showpost:", r.Form["name"][0])
	filePath := config.GConfig.PostPath + "/" + r.Form["name"][0]
	postID := zdata.GetPostIDFromPath(filePath)
	indexInfo := zdata.AllPostData[postID]
	indexInfo.PostReadNum += 1
	zdata.PageShow.PageTitle = indexInfo.PostTitle
	zdata.PageShow.PageDate = indexInfo.PostDate
	zdata.PageShow.PageContent = indexInfo.PostContent
	zdata.PageShow.PageReadNum = indexInfo.PostReadNum
	zdata.PageShow.PageCommentNum = indexInfo.PostCommentNum
	zdata.PageShow.PageComments = indexInfo.PostComments
	zdata.PageShow.BlogVersion = zversion.Ver
	zdata.AllPostData[postID] = indexInfo
	indexData := zdata.AllIndexData[postID]
	indexData.PostReadNum = indexInfo.PostReadNum
	zdata.AllIndexData[postID] = indexData
	//fmt.Println(indexInfo.PostCommentNum)
	t.Execute(w, zdata.PageShow)
}

func upcomment(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/upcomment.html")
	r.ParseForm()
	name := r.Form["name"][0]
	comment := r.Form["comment"][0]
	postPath := config.GConfig.PostPath + "/" + r.Form["postname"][0]
	postID := zdata.GetPostIDFromPath(postPath)
	dataInfo := zdata.AllPostData[postID]
	newComment := zdata.CommentStruct{
		CommentDate:     ztime.GetCurTime(ztime.DAT_MILL),
		CommentUserID:   134562,
		CommentUserName: name,
		CommentContent:  comment,
	}
	dataInfo.PostComments = append(dataInfo.PostComments, newComment)
	dataInfo.PostCommentNum = len(dataInfo.PostComments)
	zdata.AllPostData[postID] = dataInfo
	commentPath := config.GConfig.PostPath + "/" + postID + "/comment.cm"
	file.SaveComment(commentPath, dataInfo.PostComments)
	//fmt.Println(name, "::::::", comment)
	t.Execute(w, "upload comment succeed!")
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/admin.html")
	r.ParseForm()
	var iCurPage int = 0
	var err error
	// the para of page may null, check it before use
	if r.Form["page"] == nil {
		iCurPage = 1
	} else {
		iCurPage, err = strconv.Atoi(r.Form["page"][0])
		if err != nil {
			iCurPage = 1
		}
		if iCurPage <= 0 {
			iCurPage = 1
		} else if iCurPage >= zdata.IndexPage.MaxPage {
			iCurPage = zdata.IndexPage.MaxPage
		}
	}
	zdata.SetCurIndexPageShow(iCurPage)
	t.Execute(w, zdata.IndexPage)
}

func deletePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/delete.html")
	r.ParseForm()
	postID := r.Form["PostID"][0]
	var ok bool = false
	_, ok = zdata.AllPostData[postID]
	if ok {
		delete(zdata.AllPostData, postID)
	}
	_, ok = zdata.AllIndexData[postID]
	if ok {
		delete(zdata.AllIndexData, postID)
		/*
			for k, v := range zdata.AllIndexData {
				fmt.Println(k, "|||||", v)
			}
		*/
		zdata.RefreshIndexShow(zdata.AllPostData)
	}
	fmt.Println("deletePage:", postID)
	t.Execute(w, "Delete Success")
}

// temporary solution
//func getShowDownJS(w http.ResponseWriter, r *http.Request) {
//	fileContent, _ := file.ReadFile("./html/showdown.min.js")
//	w.Write([]byte(fileContent))
//}

func InitRouter() error {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/showpost", showpost)
	http.HandleFunc("/upcomment", upcomment)
	//http.HandleFunc("/showdown.min.js", getShowDownJS)
	http.HandleFunc("/admin", adminPage)
	http.HandleFunc("/delete", deletePage)

	// init static file service
	files := http.FileServer(http.Dir("./public/"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	return nil
}
