package zdata

type IndexPageStruct struct {
	WebTitle    string
	IndexData   map[string]IndexStruct
	BlogVersion string
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
