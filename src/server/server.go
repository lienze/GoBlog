package server

import (
	"GoBlog/src/config"
	"GoBlog/src/file"
	"GoBlog/src/router"
	"fmt"
	"github.com/lienze/go2db/dao"
	"net/http"
	"runtime"
)

func NewServer() error {
	var err error
	config.InitConfig()
	//fmt.Printf("%s:%d\n", config.GConfig.DB.Server, config.GConfig.DB.Port)
	addr4Server := fmt.Sprintf("%s:%d", config.GConfig.DB.Server, config.GConfig.DB.Port)
	if runtime.GOOS == "darwin" {
		fmt.Println("darwin platform")
		addr4Server = fmt.Sprintf("127.0.0.1:8080")
	}
	server := http.Server{
		//Addr: "10.0.2.15:8080",
		Addr: addr4Server,
	}
	fmt.Println("Listen:", addr4Server)

	err = router.InitRouter()
	if err != nil {
		return err
	}

	dao.InitDB("mytest")

	var mapFiles map[string]string
	mapFiles, err = file.InitFiles(config.GConfig.PostPath)
	if err != nil {
		return err
	}
	for _, val := range mapFiles {
		//fmt.Println(key, " ", val)
		router.ContentShow = append(router.ContentShow, val)
	}

	fmt.Println("GoBlog is running...")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
