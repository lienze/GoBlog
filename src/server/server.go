package server

import (
	"GoBlog/src/file"
	"GoBlog/src/router"
	"fmt"
	"github.com/lienze/go2db/dao"
	"net/http"
)

func NewServer() error {
	var err error
	server := http.Server{
		Addr: "10.0.2.15:8080",
	}
	err = router.InitRouter()
	if err != nil {
		return err
	}
	dao.InitDB("mytest")
	var mapFiles map[string]string
	mapFiles, err = file.InitFiles()
	if err != nil {
		return err
	}
	fmt.Println(len(mapFiles))
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
