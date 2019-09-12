package router

import (
	"GoBlog/src/cache"
	"GoBlog/src/config"
	"GoBlog/src/file"
	"GoBlog/src/log"
	"GoBlog/src/zdata"
	"GoBlog/src/zsession"
	"GoBlog/src/ztime"
	"GoBlog/src/zversion"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	if CheckCookie(w, r) == false {
		zdata.IndexPage.BooLogin = false
	} else {
		zdata.IndexPage.BooLogin = true
	}
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
	// NOTE: if using cache system, we are now going to record the user ip and
	// count the number of element in the cache,otherwise we just add the
	// PostReadNum value.
	if config.GConfig.Cache.Enable == true {
		//cache.AddKeyValue()
		rIdx := strings.LastIndex(r.RemoteAddr, ":")
		cache.AddSetKeyValue(postID, r.RemoteAddr[:rIdx])
		var err error
		indexInfo.PostReadNum, err = cache.GetSetSize(postID)
		//cache.PrintSet(postID)
		if err != nil {
			log.Error("cache error, see GetSetSize")
		}
	} else {
		indexInfo.PostReadNum += 1
	}
	zdata.PageShow.PageTitle = indexInfo.PostTitle
	zdata.PageShow.PageDate = indexInfo.PostDate
	zdata.PageShow.PageContent = indexInfo.PostContent
	zdata.PageShow.PageReadNum = indexInfo.PostReadNum
	zdata.PageShow.PageCommentNum = indexInfo.PostCommentNum
	zdata.PageShow.PageComments = indexInfo.PostComments
	zdata.PageShow.BlogVersion = zversion.GetVersion()
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
		CommentDate:     ztime.GetCurTime(ztime.D2AT_MILL),
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
	if CheckCookieRedirect(w, r) == false {
		return
	}
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

func newBlog(w http.ResponseWriter, r *http.Request) {
	if CheckCookieRedirect(w, r) == false {
		return
	}
	t, _ := template.ParseFiles("html/newblog.html")
	r.ParseForm()
	a := struct {
		BlogVersion string
	}{
		BlogVersion: zversion.GetVersion(),
	}
	t.Execute(w, a)
}

func deletePage(w http.ResponseWriter, r *http.Request) {
	if CheckCookieRedirect(w, r) == false {
		return
	}
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
		zdata.RefreshIndexShow()
	}
	// try to delete the folder whatever delete data succeed
	postPath := zdata.GetPostPathFromID(postID)
	//fmt.Println("deletePage:", postID)
	if err := file.RemoveFolder(postPath); err != nil {
		log.Error("delete Folder error:" + err.Error())
		t.Execute(w, "Delete Error:"+err.Error())
		return
	}
	file.SaveIndexFile(config.GConfig.PostPath+"/"+"idx.dat", zdata.AllIndexData)
	t.Execute(w, "Delete Success")
}

func savePost(w http.ResponseWriter, r *http.Request) {
	if CheckCookieRedirect(w, r) == false {
		return
	}
	t, _ := template.ParseFiles("html/savepost.html")
	r.ParseForm()
	postID := r.Form["PostID"][0]
	postTitle := r.Form["Title"][0]
	postProfile := r.Form["Profile"][0]
	postContent := r.Form["Content"][0]
	//fmt.Println("savePost:", postID)
	//fmt.Println("savePost:", postContent)
	a := struct {
		InfoString  string
		BlogVersion string
	}{
		InfoString:  "Save Succeed!",
		BlogVersion: zversion.GetVersion(),
	}
	// make sure that the postID is available in AllPostData map
	var ok bool = false
	_, ok = zdata.AllPostData[postID]
	if ok == false {
		tmp := zdata.PostStruct{
			PostID:         postID,
			PostTitle:      postTitle,
			PostProfile:    postProfile,
			PostDate:       ztime.GetCurTime(ztime.D2AT_MILL),
			PostContent:    postContent,
			PostReadNum:    0,
			PostCommentNum: 0,
			PostComments:   nil,
		}
		zdata.AllPostData[postID] = tmp
		//fmt.Println(tmp)
	}
	_, ok = zdata.AllIndexData[postID]
	if ok == false {
		tmp := zdata.IndexStruct{
			PostID:    postID,
			PostPath:  config.GConfig.PostPath + "?name=" + postID + "/" + postID + ".md",
			PostTitle: "### " + postTitle,
			PostTitleHref: "### " + "[" + postTitle + "]" +
				"(" + "./showpost?name=" + postID + "/" + postID + ".md" + ")",
			PostProfile:    ">" + postProfile,
			PostDate:       ztime.GetCurTime(ztime.D2AT_MILL),
			PostReadNum:    0,
			PostCommentNum: 0,
		}
		zdata.AllIndexData[postID] = tmp
		zdata.RefreshIndexShow()
	}
	file.CreateFolder(config.GConfig.PostPath + "/" + postID)
	file.SaveFile(config.GConfig.PostPath+"/"+postID+"/"+postID+".md", postContent)
	t.Execute(w, a)
}

func saveModifyPost(w http.ResponseWriter, r *http.Request) {
	if CheckCookieRedirect(w, r) == false {
		return
	}
	t, _ := template.ParseFiles("html/savepost.html")
	r.ParseForm()
	postID := r.Form["PostID"][0]
	postTitle := r.Form["Title"][0]
	postProfile := r.Form["Profile"][0]
	postContent := r.Form["Content"][0]
	// make sure that the postID is available in AllPostData map
	oldData, ok1 := zdata.AllPostData[postID]
	if ok1 == true {
		tmp := zdata.PostStruct{
			PostID:         postID,
			PostTitle:      postTitle,
			PostProfile:    postProfile,
			PostDate:       oldData.PostDate,
			PostContent:    postContent,
			PostReadNum:    oldData.PostReadNum,
			PostCommentNum: oldData.PostCommentNum,
			PostComments:   oldData.PostComments,
		}
		zdata.AllPostData[postID] = tmp
		//fmt.Println(tmp)
	}
	oldIndex, ok2 := zdata.AllIndexData[postID]
	if ok2 == true {
		tmp := zdata.IndexStruct{
			PostID:    postID,
			PostPath:  config.GConfig.PostPath + "?name=" + postID + "/" + postID + ".md",
			PostTitle: "### " + postTitle,
			PostTitleHref: "### " + "[" + postTitle + "]" +
				"(" + "./showpost?name=" + postID + "/" + postID + ".md" + ")",
			PostProfile:    ">" + postProfile,
			PostDate:       oldIndex.PostDate,
			PostReadNum:    oldIndex.PostReadNum,
			PostCommentNum: oldIndex.PostCommentNum,
		}
		zdata.AllIndexData[postID] = tmp
		zdata.RefreshIndexShow()
	}
	if ok1 && ok2 {
		file.RemoveFolder(config.GConfig.PostPath + "/" + postID)
		file.CreateFolder(config.GConfig.PostPath + "/" + postID)
		file.SaveFile(config.GConfig.PostPath+"/"+postID+"/"+postID+".md", postContent)
		t.Execute(w, "Save Succeed!")
	} else {
		t.Execute(w, "Save Error!")
	}
}

func modifyPost(w http.ResponseWriter, r *http.Request) {
	if CheckCookieRedirect(w, r) == false {
		return
	}
	t, _ := template.ParseFiles("html/modifypost.html")
	r.ParseForm()
	postID := r.Form["PostID"][0]
	postData := zdata.AllPostData[postID]
	fmt.Println(postData)
	a := struct {
		PostID      string
		PostTitle   string
		PostProfile string
		PostContent string
		BlogVersion string
	}{
		PostID:      postID,
		PostContent: postData.PostContent,
		PostTitle:   postData.PostTitle[4:],
		PostProfile: postData.PostProfile[1:],
		BlogVersion: zversion.GetVersion(),
	}
	//fmt.Println("modifyPost:", postID)
	t.Execute(w, a)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/login.html")
	r.ParseForm()
	userName := r.Form["Name"][0]
	passwordID := r.Form["ID"][0]
	var showInfo string
	var loginSuccess bool
	if passwordID == config.GConfig.WebSite.PassWord && userName == "admin" {
		showInfo = "Login Succeed!"
		loginSuccess = true
		zsession.GetSessionMng().AddSession(w, r)
		//fmt.Println("loginPage:", zsession.GetSessionMng())
	} else {
		showInfo = "Login Failed!"
		loginSuccess = false
	}
	a := struct {
		InfoString   string
		LoginSuccess bool
	}{
		InfoString:   showInfo,
		LoginSuccess: loginSuccess,
	}
	//fmt.Println("modifyPost:", passwordID)
	t.Execute(w, a)
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	if CheckCookieRedirect(w, r) == false {
		return
	}
	zsession.GetSessionMng().RemoveSession(w, r)
	//fmt.Println("Logout succeed!")
	http.Redirect(w, r, "/", http.StatusFound)
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
	http.HandleFunc("/newblog", newBlog)
	http.HandleFunc("/delete", deletePage)
	http.HandleFunc("/save", savePost)
	http.HandleFunc("/modify", modifyPost)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)
	http.HandleFunc("/savemodify", saveModifyPost)

	// init static file service
	files := http.FileServer(http.Dir("./public/"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	return nil
}
