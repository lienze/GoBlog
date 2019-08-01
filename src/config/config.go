package config

import (
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	PostPath string
	Host     hostsvr
	DB       database
	FileCfg  filecfg
}

type hostsvr struct {
	Server string
	Port   int
}

type database struct {
	Enable bool
	DBName string
}

type filecfg struct {
	AutoRefresh bool
	RefreshFreq int //seconds
	IgnoreFile  []string
}

var GConfig tomlConfig

func InitConfig() {
	filePath := "./config/global.toml"
	if _, err := toml.DecodeFile(filePath, &GConfig); err != nil {
		panic(err)
	}
	//fmt.Println(GConfig)
}
