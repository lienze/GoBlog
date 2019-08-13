package zdata

import (
	"GoBlog/src/config"
	"sort"
)

type ContentStruct struct {
	ContentShow []string
	MaxPage     int
	CurPage     int
	WebTitle    string
}

type IndexPageStruct struct {
	PageTitle string
	IndexData []IndexStruct
}

type IndexStruct struct {
	PageTitle   string
	PostPath    string
	PostTitle   string
	PostProfile string
	PostDate    string
}

type PageStruct struct {
	WebTitle string
	Content  string
}

var CurPageData ContentStruct
var AllPageData ContentStruct
var PageShow PageStruct
var IndexPage IndexPageStruct

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
