package router

import (
	"GoBlog/src/zsession"
	"net/http"
)

func CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	if zsession.GetSessionMng().CheckCookie(w, r) == false {
		//fmt.Println("CheckCookie Redirect to /")
		http.Redirect(w, r, "/", http.StatusFound)
		return false
	}
	return true
}
