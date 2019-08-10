package log

import (
	"GoBlog/src/config"
	"GoBlog/src/file"
	"GoBlog/src/ztime"
	"fmt"
	"strconv"
)

var (
	bEnableLog     bool = false
	logchan        chan string
	bShowInConsole bool           = false
	arrLogType                    = [...]string{"normal", "warning", "error"}
	mapLogPath     map[int]string = make(map[int]string)
	logPath        string
)

func InitLog() {
	bEnableLog = true
	fmt.Println("Init log package")
	logchan = make(chan string, 1024)
	bShowInConsole = config.GConfig.LogCfg.ShowInConsole
	fmt.Println("Log path from config file:" + config.GConfig.LogCfg.LogPath)
	logPath = config.GConfig.LogCfg.LogPath
	for key, val := range arrLogType {
		checkAndCreateFolder(logPath+val+"/", key)
	}
	go Listen4Log()
}

func checkAndCreateFolder(path string, pos int) {
	if err := file.FolderExists(path); err != nil {
		if err = file.CreateFolder(path); err != nil {
			panic(err)
		}
	}
	mapLogPath[pos] = path
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
		rawContent := recvStr[1:]
		if iType, err = strconv.Atoi(logType); err != nil {
			iType = 0
			fmt.Printf("error logType in Listen4Log [%s]\n", logType)
		}
		if bShowInConsole {
			fmt.Printf("[%s][%s] %s\n",
				ztime.GetCurTime(ztime.DAT_MILL), arrLogType[iType], rawContent)
		}
		filePath := mapLogPath[iType] + ztime.GetCurDate(ztime.STYLE1)
		fileContent := fmt.Sprintf("[%s]%s\n", ztime.GetCurTime(ztime.T_MILL), rawContent)
		file.AddContent2File(filePath, fileContent)
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
