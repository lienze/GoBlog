package zsession

import (
	"GoBlog/src/config"
	"GoBlog/src/log"
	"net/http"
)

type SessionMng struct {
	CookieName string
	SessionMap map[string]string
}

var sessionMng *SessionMng = nil

func GetSessionMng() *SessionMng {
	if sessionMng != nil {
		return sessionMng
	}
	return initSessionMng()
}

func initSessionMng() *SessionMng {
	sessionMng = &SessionMng{
		CookieName: config.GConfig.CookieCfg.CookieName,
		SessionMap: make(map[string]string),
	}
	return sessionMng
}

func (sMng *SessionMng) AddSession(w http.ResponseWriter, r *http.Request) string {
	newSessionID := "123456"
	sMng.SessionMap[newSessionID] = newSessionID
	cookie := http.Cookie{
		Name:   sMng.CookieName,
		Value:  newSessionID,
		MaxAge: config.GConfig.CookieCfg.MaxAge, // seconds
	}
	log.Normal("AddNewSession ip:" + r.RemoteAddr)
	http.SetCookie(w, &cookie)
	return newSessionID
}

func (sMng *SessionMng) CheckCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(sMng.CookieName)
	if err != nil {
		return false
	}
	if _, ok := sMng.SessionMap[cookie.Value]; ok {
		return true
	}
	return false
}

func (sMng *SessionMng) RemoveSession(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(sMng.CookieName)
	if err != nil {
		return false
	} else {
		delete(sMng.SessionMap, cookie.Value)
		newCookie := http.Cookie{
			Name:   sMng.CookieName,
			MaxAge: -1,
		}
		http.SetCookie(w, &newCookie)
	}
	return true
}

func (sMng *SessionMng) GetSessionNum() int {
	return len(sMng.SessionMap)
}
