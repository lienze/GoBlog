package zdata

type IndexPageStruct struct {
	WebTitle     string
	AllIndexData map[string]IndexStruct
	AllIndexKey  []string
	CurIndexData []IndexStruct
	CurPage      int
	MaxPage      int
	BlogVersion  string
}

type IndexStruct struct {
	PostPath       string
	PostTitle      string
	PostTitleHref  string
	PostProfile    string
	PostDate       string
	PostReadNum    int
	PostCommentNum int
}
