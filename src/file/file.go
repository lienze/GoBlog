package file

import (
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
	for _, f := range files {
		fileFullPath := postPath + f.Name()
		if retContent, err := ReadFile(fileFullPath); err == nil {
			retMapFileContent[fileFullPath] = retContent
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
