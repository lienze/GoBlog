package log

import (
	"GoBlog/src/config"
	"GoBlog/src/file"
	"GoBlog/src/gtime"
	"fmt"
	"strconv"
)

var (
	bEnableLog     bool = false
	logchan        chan string
	bShowInConsole bool = false
	arrLogType          = [...]string{"normal", "warning", "error"}
	logPath        string
)

func InitLog() {
	bEnableLog = true
	fmt.Println("Init log package")
	logchan = make(chan string, 1024)
	bShowInConsole = config.GConfig.LogCfg.ShowInConsole
	fmt.Println("Log path from config file:" + config.GConfig.LogCfg.LogPath)
	logPath = config.GConfig.LogCfg.LogPath
	//check if the logPath exists
	if err := file.FolderExists(logPath); err != nil {
		if err = file.CreateFolder(logPath); err != nil {
			panic(err)
		}
	}
	go Listen4Log()
}

func wlog(info string) {
	logchan <- info
}

func Listen4Log() {
	var recvStr string
	var iType int = 0
	var err error
	select {
	case recvStr = <-logchan:
		logType := recvStr[0:1]
		if iType, err = strconv.Atoi(logType); err != nil {
			iType = 0
			fmt.Printf("error logType in Listen4Log [%s]\n", logType)
		}
		if bShowInConsole {
			fmt.Printf("[%s][%s] %s\n",
				gtime.GetCurTime(gtime.BASIC_MILL), arrLogType[iType], recvStr[1:])
		}
		//file.SaveFile(logPath+"firstFile", "Hello World")
		file.AddContent2File(logPath+"firstFile", "Hello World")
	}
}

func Normal(rawinfo string) {
	if !bEnableLog {
		return
	}
	wlog(NORMAL + rawinfo)
}

func Warning(rawinfo string) {
	if !bEnableLog {
		return
	}
	wlog(WARNING + rawinfo)
}

func Error(rawinfo string) {
	if !bEnableLog {
		return
	}
	wlog(ERROR + rawinfo)
}
