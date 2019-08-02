package log

import (
	"GoBlog/src/config"
	"fmt"
)

var (
	bEnableLog     bool = false
	logchan        chan string
	bShowInConsole bool = false
)

func InitLog() {
	bEnableLog = true
	fmt.Println("init log package")
	logchan = make(chan string, 1024)
	bShowInConsole = config.GConfig.LogCfg.ShowInConsole
	fmt.Println(config.GConfig.LogCfg.LogPath)
	go Listen4Log()
}

func wlog(info string) {
	logchan <- info
}

func Listen4Log() {
	var recvStr string
	select {
	case recvStr = <-logchan:
		if bShowInConsole {
			fmt.Println(recvStr)
		}
	}
}

func Normal(rawinfo string) {
	if bEnableLog {
		return
	}
	wlog(NORMAL + rawinfo)
}

func Warning(rawinfo string) {
	if bEnableLog {
		return
	}
	wlog(WARNING + rawinfo)
}

func Error(rawinfo string) {
	if bEnableLog {
		return
	}
	wlog(ERROR + rawinfo)
}
