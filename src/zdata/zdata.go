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

var PageShow PageStruct
var IndexPage IndexPageStruct

var AllIndexData map[string]IndexStruct
var AllPostData map[string]PostStruct

func RefreshIndexShow() {
	var mapkeys []string
	for k := range AllIndexData {
		mapkeys = append(mapkeys, k)
	}
	//fmt.Println(mapkeys)
	sort.Sort(sort.Reverse(sort.StringSlice(mapkeys)))
	IndexPage.WebTitle = config.GConfig.WebSite.WebTitle
	IndexPage.AllIndexKey = mapkeys
	SetCurIndexPageShow(1)
}

func SetCurIndexPageShow(iCurPage int) {
	mapkeys := IndexPage.AllIndexKey
	iPerPage := config.GConfig.PageCfg.MaxItemPerPage
	iAllDataLen := len(IndexPage.AllIndexKey)
	iMaxPage := len(mapkeys) / iPerPage
	if len(mapkeys)%iPerPage != 0 {
		iMaxPage++
	}
	if iMaxPage <= 0 {
		iMaxPage = 1
	}
	IndexPage.MaxPage = iMaxPage
	IndexPage.CurPage = iCurPage
	IndexPage.CurIndexData = make([]IndexStruct, 0)
	for i := (iCurPage - 1) * iPerPage; i < iCurPage*iPerPage && i < iAllDataLen; i++ {
		k := IndexPage.AllIndexKey[i]
		IndexPage.CurIndexData =
			append(IndexPage.CurIndexData, AllIndexData[k])
	}
}

// collect post data through IndexPage and MapFile struct
func RefreshAllPostData(mapFiles map[string]string, mapComments map[string][]CommentStruct) {
	AllPostData = make(map[string]PostStruct)
	for k, v := range mapFiles {
		comm := mapComments[k]
		comms := make([]CommentStruct, 0)
		comms = append(comms, comm...)
		indexData, ok := AllIndexData[k]
		if ok == false {
			continue
		}
		indexData.PostCommentNum = len(comms)
		AllIndexData[k] = indexData
		tmp := PostStruct{
			PostID:         k,
			PostTitle:      indexData.PostTitle,
			PostProfile:    indexData.PostProfile,
			PostDate:       indexData.PostDate,
			PostContent:    v,
			PostReadNum:    indexData.PostReadNum,
			PostCommentNum: indexData.PostCommentNum,
			PostComments:   comms,
		}
		//commentPath := GetCommentPathFromID(k)
		//fmt.Println("[RefreshAllPostData]commentPath:", commentPath)
		AllPostData[k] = tmp
	}
}

// refresh index data
func RefreshAllIndexData() {
	for k, v := range AllIndexData {
		v.PostCommentNum = len(AllPostData[k].PostComments)
		AllIndexData[k] = v
	}
}
