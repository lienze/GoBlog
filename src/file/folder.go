package file

import (
	"GoBlog/src/config"
	"os"
	"time"
)

func ScanFolder(postPath string) {
	if config.GConfig.FileCfg.AutoRefresh == true {
		freq := config.GConfig.FileCfg.RefreshFreq
		for {
			time.Sleep(time.Duration(freq) * time.Second)

		}
	}

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
