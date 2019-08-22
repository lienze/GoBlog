package file

import (
	"GoBlog/src/config"
	"GoBlog/src/zdata"
	"GoBlog/src/zversion"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	bUseFilePool bool = false
	mapFilePool  map[string]*os.File
)

// convert from struct to []byte
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func InitFiles(postPath string) (map[string]string, map[string][]zdata.CommentStruct, error) {
	//fmt.Println("InitFiles...")
	// init options
	bUseFilePool = config.GConfig.FileCfg.UseFilePool
	mapFilePool = make(map[string]*os.File)
	loadIndexData()
	return LoadFiles(postPath)
}

func LoadFiles(postPath string) (map[string]string, map[string][]zdata.CommentStruct, error) {
	fmt.Println("Start Loading Files...")
	retMapFileContent := make(map[string]string)
	retMapFileComment := make(map[string][]zdata.CommentStruct)
	readPath(postPath, &retMapFileContent, &retMapFileComment)
	fmt.Println("Load Files ok...")
	return retMapFileContent, retMapFileComment, nil
}

func readPath(postRootPath string,
	retMapFileContent *map[string]string,
	retMapFileComment *map[string][]zdata.CommentStruct) error {
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
			postID := zdata.GetPostIDFromPath(fileFullPath)
			if ext == "md" {
				if retContent, err := ReadFile(fileFullPath); err == nil {
					(*retMapFileContent)[postID] = retContent
				}
			} else if ext == "cm" {
				retSlice, _ := analyseComments(fileFullPath)
				//fmt.Println(r)
				(*retMapFileComment)[postID] = append((*retMapFileComment)[postID], retSlice...)
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

func SaveComment(filename string, content []zdata.CommentStruct) {
	var writeData string
	for _, v := range content {
		//fmt.Println(k, " ", v)
		writeData += v.CommentDate + "@" + strconv.FormatInt(v.CommentUserID, 10) + "@" +
			v.CommentUserName + "@" + v.CommentContent + "\n"
		// convert CommentStruct to string
		/*
			iLen := unsafe.Sizeof(v)
			bBytes := &SliceMock{
				addr: uintptr(unsafe.Pointer(&v)),
				cap:  int(iLen),
				len:  int(iLen),
			}
			data := *(*[]byte)(unsafe.Pointer(bBytes))
		*/
	}
	SaveFile(filename, writeData)
}

func SaveIndexFile(filePath string, content map[string]zdata.IndexStruct) {
	var writeData string
	for _, v := range content {
		//fmt.Println("SaveIndexFile:", k, "!!!!!!", v)
		writeData += v.PostID + "@" + v.PostTitle[4:] + "@" +
			v.PostProfile[1:] + "@" + v.PostDate + "@" +
			strconv.Itoa(v.PostReadNum) + "@" +
			strconv.Itoa(v.PostCommentNum) + "\n"
	}
	//fmt.Println(writeData)
	SaveFile(filePath, writeData)
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
			sList := strings.Split(filename, "/")
			// example: filename is "./log/normal/20190804"
			logType := sList[2]
			for key, _ := range mapFilePool {
				if strings.Contains(key, logType) {
					mapFilePool[key].Close()
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
	fileObj, err := os.OpenFile(config.GConfig.PostPath+"/"+"idx.dat", os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	buf := bufio.NewReader(fileObj)
	zdata.AllIndexData = make(map[string]zdata.IndexStruct)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			//fmt.Println("loadIndexData:", line)
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
				PostID:    sList[0],
				PostPath:  config.GConfig.PostPath + "?name=" + sList[0] + "/" + sList[0] + ".md",
				PostTitle: "### " + sList[1],
				PostTitleHref: "### " + "[" + sList[1] + "]" +
					"(" + "./showpost?name=" + sList[0] + "/" + sList[0] + ".md" + ")",
				PostProfile:    ">" + sList[2],
				PostDate:       sList[3],
				PostReadNum:    postReadNum,
				PostCommentNum: postCommentNum,
			}
			//k := zdata.GetPostIDFromPath(config.GConfig.PostPath + "/" + sList[0])
			zdata.AllIndexData[sList[0]] = tmp
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

func analyseComments(commentPath string) ([]zdata.CommentStruct, error) {
	fmt.Println("analyseComments begin...")
	fileObj, err := os.OpenFile(commentPath, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer fileObj.Close()
	buf := bufio.NewReader(fileObj)
	var ret []zdata.CommentStruct
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			sList := strings.Split(line, "@")
			/*parseTime, parseErr := time.Parse("2006-01-02 15:04:05.000", sList[0])
			if parseErr == nil {
				fmt.Println("parseTime:", parseTime)
			}*/
			commentUserID, errconv := strconv.ParseInt(sList[1], 10, 64)
			if errconv != nil {
				commentUserID = -1
			}
			tmp := zdata.CommentStruct{
				CommentDate:     sList[0],
				CommentUserID:   commentUserID,
				CommentUserName: sList[2],
				CommentContent:  sList[3],
			}
			//fmt.Println(tmp)
			ret = append(ret, tmp)
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}
	fmt.Println("analyseComments end...")
	return ret, nil
}
