package zdata

type PageStruct struct {
	WebTitle       string
	PageTitle      string
	PageDate       string
	PageContent    string
	PageReadNum    int
	PageCommentNum int
	PageComments   []CommentStruct
	BlogVersion    string
}
