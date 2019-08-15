package file

import (
	"GoBlog/src/config"
	"GoBlog/src/zdata"
	"GoBlog/src/ztime"
	"GoBlog/src/zversion"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	bUseFilePool bool = false
	mapFilePool  map[string]*os.File
)

func InitFiles(postPath string) (map[string]string, map[string]zdata.CommentStruct, error) {
	//fmt.Println("InitFiles...")
	// init options
	bUseFilePool = config.GConfig.FileCfg.UseFilePool
	mapFilePool = make(map[string]*os.File)
	loadIndexData()
	return LoadFiles(postPath)
}

func LoadFiles(postPath string) (map[string]string, map[string]zdata.CommentStruct, error) {
	fmt.Println("Start Loading Files...")
	retMapFileContent := make(map[string]string)
	retMapFileComment := make(map[string]zdata.CommentStruct)
	readPath(postPath, &retMapFileContent, &retMapFileComment)
	fmt.Println("Load Files ok...")
	return retMapFileContent, retMapFileComment, nil
}

func readPath(postRootPath string,
	retMapFileContent *map[string]string,
	retMapFileComment *map[string]zdata.CommentStruct) error {
	files, errDir := ioutil.ReadDir(postRootPath)
	if errDir != nil {
		return errDir
	}
	//fmt.Println(config.GConfig.FileCfg)
	includeFileArr := config.GConfig.FileCfg.IncludeFile
	for _, f := range files {
		fileFullPath := postRootPath + f.Name()
		if f.IsDir() {
			readPath(fileFullPath+"/", retMapFileContent, retMapFileComment)
			continue
		}
		// check file ext
		ext := getFileExt(fileFullPath)
		bIgnore := true
		for _, val := range includeFileArr {
			if ext == val {
				bIgnore = false
				break
			}
		}
		if !bIgnore {
			if retContent, err := ReadFile(fileFullPath); err == nil {
				postID := zdata.GetPostIDFromPath(fileFullPath)
				if ext == "md"{
					(*retMapFileContent)[postID] = retContent
				} else if ext == "cm"{
					sList := strings.Split(retContent, "@")
					commentUserID, errconv := strconv.ParseInt(sList[0], 10, 64)
					if errconv != nil {
						commentUserID = -1
					}
					tmp := zdata.CommentStruct{
						CommentDate:     time.Now(),
						CommentDateShow: ztime.GetCurTime(ztime.DAT),
						CommentUserID:   commentUserID,
						CommentUserName: sList[1],
						CommentContent:  sList[2],
					}
					fmt.Println(tmp)
					(*retMapFileComment)[postID] = tmp
				}
			}
		}
	}
	return nil
}

func ReadFile(name string) (string, error) {
	fmt.Println("Start ReadFile", name)
	if contents, err := ioutil.ReadFile(name); err == nil {
		result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("content:", string(result))
		return result, nil
	} else {
		fmt.Println(err)
		return "", err
	}
}

func SaveFile(filename string, content string) error {
	fmt.Println("Start SaveFile", filename)
	data := []byte(content)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func FileExist(filePath string) {
	_, err := os.Stat(filePath)
	if err == nil {
		fmt.Println("FileExist:", filePath)
	} else {
		fmt.Println("FileNotExist:", filePath)
	}
}

func AddContent2File(filename string, content string) error {
	//fmt.Println("Start Add Content 2 File")
	bNewFile := true
	var fileObj *os.File = nil
	var bKeyExist bool = false
	var err error
	if bUseFilePool {
		if fileObj, bKeyExist = mapFilePool[filename]; bKeyExist == true {
			bNewFile = false
		}
	}
	if bNewFile == true {
		fileObj, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
	}
	if bUseFilePool == true {
		// do not close file obj when use file pool option
		if bNewFile == true {
			// close the same type log before
			slist := strings.Split(filename, "/")
			// example: filename is "./log/normal/20190804"
			logType := slist[2]
			for key, _ := range mapFilePool {
				if strings.Contains(key, logType) {
					delete(mapFilePool, key)
				}
			}
			mapFilePool[filename] = fileObj
			fmt.Printf("added fileObj to mapFilePool:[%s] %s\n", logType, filename)
		}
	} else {
		defer fileObj.Close()
	}
	if _, err = fileObj.WriteString(content); err != nil {
		return err
	}
	return nil
}

func getFileExt(filename string) string {
	idx := strings.LastIndex(filename, ".")
	//fmt.Println(filename[idx:])
	return string(filename[idx+1:])
}

func loadIndexData() error {
	fmt.Println("loadIndexData begin...")
	fileObj, err := os.OpenFile(config.GConfig.PostPath+"idx.dat", os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	buf := bufio.NewReader(fileObj)
	zdata.IndexPage.IndexData = make(map[string]zdata.IndexStruct)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			//fmt.Println(line)
			sList := strings.Split(line, "@")
			postReadNum, errconv := strconv.Atoi(sList[4])
			if errconv != nil {
				postReadNum = -1
			}
			postCommentNum, errconv := strconv.Atoi(sList[5])
			if errconv != nil {
				postCommentNum = -1
			}
			tmp := zdata.IndexStruct{
				PostPath:  "./post?name=" + sList[0],
				PostTitle: "### " + sList[1],
				PostTitleHref: "### " + "[" + sList[1] + "]" +
					"(" + "./showpost?name=" + sList[0] + ")",
				PostProfile:    ">" + sList[2],
				PostDate:       sList[3],
				PostReadNum:    postReadNum,
				PostCommentNum: postCommentNum,
			}
			k := zdata.GetPostIDFromPath(config.GConfig.PostPath + sList[0])
			zdata.IndexPage.IndexData[k] = tmp
			//fmt.Println(tmp)
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read idx.dat ok!")
				break
			} else {
				fmt.Println("Read file error", err)
				return err
			}
		}
	}
	zdata.IndexPage.WebTitle = config.GConfig.WebSite.WebTitle
	zdata.IndexPage.BlogVersion = zversion.Ver
	fmt.Println("loadIndexData end...")
	return nil
}

func loadComments() {

}
