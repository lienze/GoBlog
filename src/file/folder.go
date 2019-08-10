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
	var err error
	if config.GConfig.FileCfg.AutoRefresh == true {
		freq := config.GConfig.FileCfg.RefreshFreq
		for {
			time.Sleep(time.Duration(freq) * time.Second)
			filesInfo, errDir := ReadFolder(postPath, config.GConfig.FileCfg.IgnoreFile)
			if errDir != nil {
				fmt.Println(errDir)
			}
			if len(filesInfo) != len(MapFiles) {
				// files in postPath folder have been changed
				// update right now
				// FIXME: there seem to have a bug then we add a file and remove
				//        a file at the same time, files have been changed but
				//        the process could not find the difference, so no refresh
				//        happened.We may use MD5 to compare.
				MapFiles = make(map[string]string)
				MapFiles, err = LoadFiles(config.GConfig.PostPath)
				if err == nil {
					fmt.Println("Finished LoadFiles")
					zdata.RefreshContentShow(MapFiles)
				} else {
					fmt.Println("ScanFolder...", err)
				}
			}
			//fmt.Println("ScanFolder:", len(filesInfo), len(MapFiles))
		}
	}
}

func ReadFolder(postPath string, ignoreExt []string) ([]os.FileInfo, error) {
	var retFilesInfo []os.FileInfo
	filesInfo, errDir := ioutil.ReadDir(postPath)
	if len(ignoreExt) > 0 {
		for _, f := range filesInfo {
			fileFullPath := postPath + f.Name()
			//check ignore file
			ext := getFileExt(fileFullPath)
			bIgnore := false
			for _, val := range ignoreExt {
				if ext == val {
					bIgnore = true
					break
				}
			}
			if !bIgnore {
				retFilesInfo = append(retFilesInfo, f)
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
