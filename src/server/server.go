package server

import (
	"GoBlog/src/cache"
	"GoBlog/src/config"
	"GoBlog/src/db"
	"GoBlog/src/file"
	"GoBlog/src/log"
	"GoBlog/src/router"
	"GoBlog/src/zdata"
	"GoBlog/src/zversion"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func NewServer() error {
	var err error
	config.InitConfig()
	//fmt.Printf("%s:%d\n", config.GConfig.DB.Server, config.GConfig.DB.Port)
	addr4Server := fmt.Sprintf("%s:%d",
		config.GConfig.Host.Server,
		config.GConfig.Host.Port)
	fmt.Println("GoBlog version:", zversion.GetVersion())
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

	if config.GConfig.DB.Enable == true {
		if err := db.InitDB(); err != nil {
			log.Error(err.Error())
			panic(err)
		}
	}

	if config.GConfig.Cache.Enable == true {
		if err := cache.InitCache(config.GConfig.Cache.CacheType); err != nil {
			log.Error(err.Error())
			panic(err)
		}
	}

	if config.GConfig.LogCfg.Enable == true {
		log.InitLog()
	}

	var mapFiles map[string]string
	var mapComments map[string][]zdata.CommentStruct
	mapFiles, mapComments, err = file.InitFiles()
	if err != nil {
		return err
	}

	zdata.InitAllPostData(&mapFiles, &mapComments)

	zdata.RefreshIndexShow()
	// new gorountine for scanning folder that we could refresh page
	// when there is new post appear
	//go file.ScanFolder(config.GConfig.PostPath + "/")

	// catch signal
	go HandleSignal()

	log.Normal("This is a normal info")
	log.Warning("This is a warning info")
	log.Error("This is an error info")

	fmt.Printf("GoBlog is running...http://%s\n", addr4Server)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func HandleSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-c:
		fmt.Printf("\nGet [%s] signal\n", sig)
		handSignal(sig)
		os.Exit(0)
	}
}

func handSignal(sig os.Signal) {
	switch sig {
	case syscall.SIGTERM, os.Interrupt:
		fmt.Println("hand", sig)
		fmt.Println("SaveIndexFile number:", len(zdata.AllIndexData))
		file.SaveIndexFile(config.GConfig.PostPath+"/"+"idx.dat", zdata.AllIndexData)
	default:
		fmt.Printf("hand [%s]\n", sig)
	}
}
