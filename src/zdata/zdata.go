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

var AllPostData map[string]PostStruct
var AllCommentData map[string]CommentStruct

func RefreshIndexShow(mapFiles map[string]PostStruct) {
	var mapkeys []string
	for k := range IndexPage.AllIndexData {
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
	for _, val := range mapkeys {
		//fmt.Println(key, " ", val)
		IndexPage.CurIndexData = append(IndexPage.CurIndexData, IndexPage.AllIndexData[val])
	}
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
			append(IndexPage.CurIndexData, IndexPage.AllIndexData[k])
	}
}

// collect post data through IndexPage and MapFile struct
func RefreshAllPostData(mapFiles map[string]string, mapComments map[string][]CommentStruct) {
	AllPostData = make(map[string]PostStruct)
	for k, v := range mapFiles {
		indexData := IndexPage.AllIndexData[k]
		comm := mapComments[k]
		comms := make([]CommentStruct, 0)
		comms = append(comms, comm...)
		indexData.PostCommentNum = len(comms)
		IndexPage.AllIndexData[k] = indexData
		tmp := PostStruct{
			PostPath:       k,
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
