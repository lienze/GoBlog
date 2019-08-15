package zdata

import "time"

type CommentStruct struct {
	CommentDate     time.Time
	CommentUserID   uint64
	CommentUserName string
	CommentConent   string
}

type PostStruct struct {
	PostPath       string
	PostTitle      string
	PostProfile    string
	PostDate       string
	PostContent    string
	PostReadNum    int
	PostCommentNum int
	PostComments   []CommentStruct
}
