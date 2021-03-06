package zdata

import (
	"GoBlog/src/config"
	"strings"
)

type CommentStruct struct {
	CommentDate      string
	CommentUserID    int64
	CommentUserName  string
	CommentUserEmail string
	CommentContent   string
}

type PostStruct struct {
	PostID         string
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

func GetPostRootPathFromID(postID string) string {
	return config.GConfig.PostPath + "/" + postID + "/"
}

func GetPostPathFromID(postID string) string {
	return config.GConfig.PostPath + "/" + postID + "/" + postID + ".md"
}

func GetCommentPathFromID(postID string) string {
	return config.GConfig.PostPath + "/" + postID + "/comment.cm"
}
