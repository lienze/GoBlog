package zdata

type IndexPageStruct struct {
	WebTitle     string
	AllIndexKey  []string
	CurIndexData []IndexStruct
	CurPage      int
	MaxPage      int
	BooLogin     bool
	BlogVersion  string
}

type IndexStruct struct {
	PostID         string
	PostPath       string
	PostTitle      string
	PostTitleHref  string
	PostProfile    string
	PostDate       string
	PostReadNum    int
	PostCommentNum int
}
