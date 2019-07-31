package config

import (
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	PostPath string
	Host     hostsvr
	DB       database
}

type hostsvr struct {
	Server string
	Port   int
}

type database struct {
	DBAble bool
	DBName string
}

var GConfig tomlConfig

func InitConfig() {
	filePath := "./config/global.toml"
	if _, err := toml.DecodeFile(filePath, &GConfig); err != nil {
		panic(err)
	}
	//fmt.Println(GConfig)
}
