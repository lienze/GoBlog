package server

import (
	"GoBlog/src/config"
	"GoBlog/src/file"
	"GoBlog/src/router"
	"fmt"
	"github.com/lienze/go2db/dao"
	"net/http"
)

func NewServer() error {
	var err error
	config.InitConfig()
	//fmt.Printf("%s:%d\n", config.GConfig.DB.Server, config.GConfig.DB.Port)
	addr4Server := fmt.Sprintf("%s:%d", config.GConfig.DB.Server, config.GConfig.DB.Port)
	server := http.Server{
		//Addr: "10.0.2.15:8080",
		Addr: addr4Server,
	}
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
	for key, val := range mapFiles {
		fmt.Println(key, " ", val)
	}
	fmt.Println("GoBlog is running...")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
