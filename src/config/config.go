package config

import (
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	PostPath string
	DB       database
}

type database struct {
	Server string
	Port   int
}

var GConfig tomlConfig

func InitConfig() {
	filePath := "./config/global.toml"
	if _, err := toml.DecodeFile(filePath, &GConfig); err != nil {
		panic(err)
	}
	//fmt.Println(GConfig)
}
