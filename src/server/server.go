package server

import (
	"GoBlog/src/config"
	"GoBlog/src/db"
	"GoBlog/src/file"
	"GoBlog/src/log"
	"GoBlog/src/router"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"syscall"
)

func NewServer() error {
	var err error
	config.InitConfig()
	//fmt.Printf("%s:%d\n", config.GConfig.DB.Server, config.GConfig.DB.Port)
	addr4Server := fmt.Sprintf("%s:%d",
		config.GConfig.Host.Server,
		config.GConfig.Host.Port)

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
			panic(err)
		}
	}

	fmt.Println(config.GConfig.LogCfg)
	if config.GConfig.LogCfg.Enable == true {
		log.InitLog()
	}

	var mapFiles map[string]string
	mapFiles, err = file.InitFiles(config.GConfig.PostPath)
	if err != nil {
		return err
	}
	// new gorountine for scanning folder when there is new posts appear
	go file.ScanFolder(config.GConfig.PostPath)
	var mapkeys []string
	for k := range mapFiles {
		mapkeys = append(mapkeys, k)
	}
	//fmt.Println(mapkeys)
	sort.Sort(sort.Reverse(sort.StringSlice(mapkeys)))
	for _, val := range mapkeys {
		//fmt.Println(key, " ", val)
		router.ContentShow = append(router.ContentShow, mapFiles[val])
	}

	//catch signal
	go HandleSignal()

	log.Normal("Hello World!")

	fmt.Println("GoBlog is running...")
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
		fmt.Printf("Get [%s] signal\n", sig)
		handSignal(sig)
		os.Exit(0)
	}
}

func handSignal(sig os.Signal) {
	switch sig {
	case syscall.SIGTERM:
		fmt.Println("hand SIGTERM")
	case os.Interrupt:
		fmt.Println("hand Interrupt")
	default:
		fmt.Printf("hand [%s]\n", sig)
	}
}
