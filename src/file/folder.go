package file

import (
	"GoBlog/src/config"
	"GoBlog/src/zdata"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func ScanFolder(postPath string) {
	if config.GConfig.FileCfg.AutoRefresh == true {
		freq := config.GConfig.FileCfg.RefreshFreq
		for {
			time.Sleep(time.Duration(freq) * time.Second)
			filesInfo, errDir := ReadFolder(postPath, config.GConfig.FileCfg.IncludeFile)
			if errDir != nil {
				fmt.Println(errDir)
			}
			if len(filesInfo) != len(zdata.AllPostData) {
				// files in postPath folder have been changed
				// update right now
				// FIXME: there seem to have a bug then we add a file and remove
				//        a file at the same time, files have been changed but
				//        the process could not find the difference, so no refresh
				//        happened.We may use MD5 to compare.
				mapFiles, mapComments, err := LoadFiles(config.GConfig.PostPath)
				if err == nil {
					fmt.Println("Finished LoadFiles")
					zdata.RefreshIndexShow(zdata.AllPostData)
					zdata.RefreshAllPostData(mapFiles, mapComments)
				} else {
					fmt.Println("ScanFolder...", err)
				}
				fmt.Println("ScanFolder:", len(filesInfo), len(mapFiles))
			}
		}
	}
}

func ReadFolder(postPath string, includeExt []string) ([]os.FileInfo, error) {
	var retFilesInfo []os.FileInfo
	filesInfo, errDir := ioutil.ReadDir(postPath)
	if len(includeExt) > 0 {
		for _, f := range filesInfo {
			fileFullPath := postPath + "/" + f.Name()
			if f.IsDir() {
				fList, err := ReadFolder(fileFullPath, includeExt)
				if err == nil {
					retFilesInfo = append(retFilesInfo, fList...)
				}
				continue
			}
			//check ignore file
			ext := getFileExt(fileFullPath)
			bIgnore := true
			for _, val := range includeExt {
				if ext == val {
					bIgnore = false
					break
				}
			}
			if !bIgnore {
				if ext == "md" {
					retFilesInfo = append(retFilesInfo, f)
				}
			}
		}
	}
	return retFilesInfo, errDir
}

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
