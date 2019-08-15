package zdata

import (
	"GoBlog/src/config"
	"strings"
	"time"
)

type CommentStruct struct {
	CommentDate     time.Time
	CommentDateShow string
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

func GetPostIDFromPath(fileFullPath string) string {
	idx1 := strings.LastIndex(fileFullPath, "/")
	idx2 := strings.LastIndex(fileFullPath[:idx1], "/")
	postID := fileFullPath[idx2+1 : idx1]
	return postID
}

func GetPostPathFromID(postID string) string {
	return config.GConfig.PostPath + postID + "/"
}

func GetCommentPathFromID(postID string) string {
	return config.GConfig.PostPath + postID + "/comment.dat"
}
