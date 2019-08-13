package file

import (
	"GoBlog/src/config"
	"GoBlog/src/zdata"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	bUseFilePool bool = false
	mapFilePool  map[string]*os.File
	MapFiles     map[string]string // the whole data
)

func InitFiles(postPath string) (map[string]string, error) {
	//fmt.Println("InitFiles...")
	// init options
	bUseFilePool = config.GConfig.FileCfg.UseFilePool
	MapFiles = make(map[string]string)
	mapFilePool = make(map[string]*os.File)
	loadIndexData()
	return LoadFiles(postPath)
}

func LoadFiles(postPath string) (map[string]string, error) {
	fmt.Println("Start Loading Files...")
	retMapFileContent := make(map[string]string)
	readPath(postPath, &retMapFileContent)
	fmt.Println("Load Files ok...")
	return retMapFileContent, nil
}

func readPath(postRootPath string, retMapFileContent *map[string]string) error {
	files, errDir := ioutil.ReadDir(postRootPath)
	if errDir != nil {
		return errDir
	}
	//fmt.Println(config.GConfig.FileCfg)
	includeFileArr := config.GConfig.FileCfg.IncludeFile
	for _, f := range files {
		fileFullPath := postRootPath + f.Name()
		if f.IsDir() {
			readPath(fileFullPath+"/", retMapFileContent)
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
				(*retMapFileContent)[fileFullPath] = retContent
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
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			//fmt.Println(line)
			slist := strings.Split(line, "@")
			tmp := zdata.IndexStruct{
				PostPath:  slist[0],
				PostTitle: slist[1] + "(" + slist[0] + ")",
				//PostTitle:   slist[1] + "(http://www.baidu.com)",
				PostProfile: slist[2],
				PostDate:    slist[3],
			}
			zdata.IndexData = append(zdata.IndexData, tmp)
			//fmt.Println(zdata.IndexData)
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
	fmt.Println("loadIndexData end...")
	return nil
}
