package file

import (
	"GoBlog/src/config"
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
