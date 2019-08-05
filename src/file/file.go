package file

import (
	"GoBlog/src/config"
	"fmt"
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
	return LoadFiles(postPath)
}

func LoadFiles(postPath string) (map[string]string, error) {
	retMapFileContent := make(map[string]string)
	files, errDir := ioutil.ReadDir(postPath)
	if errDir != nil {
		return nil, errDir
	}
	//fmt.Println(config.GConfig.FileCfg)
	ignoreFileArr := config.GConfig.FileCfg.IgnoreFile
	for _, f := range files {
		fileFullPath := postPath + f.Name()
		//check ignore file
		ext := getFileExt(fileFullPath)
		bIgnore := false
		for _, val := range ignoreFileArr {
			if ext == val {
				bIgnore = true
				break
			}
		}
		if !bIgnore {
			if retContent, err := ReadFile(fileFullPath); err == nil {
				retMapFileContent[fileFullPath] = retContent
			}
		}
	}
	return retMapFileContent, nil
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
			fmt.Printf("mapFilePool:%s %s\n", logType, filename)
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
