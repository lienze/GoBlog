package log

import (
	"os"
)

func FolderExists(folderPath string) error {
	var err error
	if _, err = os.Stat(folderPath); err == nil {
		return nil
	} else {
		return err
	}
}

func CreateFolder(folderPath string) error {
	err := os.MkdirAll(folderPath, 0711)
	return err
}

func AddContent2File(filename string, content string) error {
	//fmt.Println("Start Add Content 2 File")
	var fileObj *os.File = nil
	var err error

	fileObj, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer fileObj.Close()
	if _, err = fileObj.WriteString(content); err != nil {
		return err
	}
	return nil
}
