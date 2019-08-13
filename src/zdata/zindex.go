package zdata

type IndexPageStruct struct {
	PageTitle string
	IndexData []IndexStruct
}

type IndexStruct struct {
	PostPath       string
	PostTitle      string
	PostProfile    string
	PostDate       string
	PostReadNum    int
	PostCommentNum int
}
