package file

import (
	"GoBlog/src/config"
	"fmt"
	"io/ioutil"
	"strings"
)

func InitFiles(postPath string) (map[string]string, error) {
	//fmt.Println("InitFiles...")
	retMapFileContent := make(map[string]string)
	files, errDir := ioutil.ReadDir(postPath)
	if errDir != nil {
		return nil, errDir
	}
	//fmt.Println(config.GConfig.FileCfg.IgnoreFile)
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

// private function
func getFileExt(filename string) string {
	idx := strings.LastIndex(filename, ".")
	//fmt.Println(filename[idx:])
	return string(filename[idx+1:])
}
